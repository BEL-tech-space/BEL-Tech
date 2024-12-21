package http

import (
	"beltech/service"
	"bytes"
	"context"
	"errors"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"io"
	"net/http"
	"strconv"

	agentPages "beltech/frontend/pages/agent"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/nfnt/resize"
	"github.com/skip2/go-qrcode"
)

func (a *App) collectPage(w http.ResponseWriter, r *http.Request) {

	agentID := "devMode"
	if a.config.Build == productionEnv {
		agentID = r.Context().Value(contextKeyAgentID).(string)
	}

	if err := agentPages.CollectPage(agentID).Render(context.Background(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *App) wasteForm(w http.ResponseWriter, r *http.Request) {
	agentID := "devMode"
	if a.config.Build == productionEnv {
		agentID = r.Context().Value(contextKeyAgentID).(string)
	}

	if err := agentPages.WasteForm(agentID).Render(context.Background(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *App) generateQRCode(w http.ResponseWriter, r *http.Request) {

	var (
		agentUUID uuid.UUID = uuid.Nil
		err       error
	)

	if a.config.Build == productionEnv {
		agentID := chi.URLParam(r, "agentID")
		if agentID == "" {
			http.Error(w, "agentID not specified", http.StatusBadRequest)
			return
		}

		if agentUUID, err = uuid.Parse(agentID); err != nil {
			http.Error(w, "invalid agentID", http.StatusBadRequest)
			return
		}
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	glass, paper, metal, plastics, err := getFormattedData(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var wasteID uuid.UUID
	p := service.AddNewWasteFromCollectorParams{
		GlassWeighed:    glass,
		MetalWeighed:    metal,
		PlasticsWeighed: plastics,
		PaperWeighed:    paper,
		FunnyName:       gofakeit.Name(),
		AgentID:         agentUUID,
	}
	if a.config.Build == productionEnv {
		// if wasteID, err = a.service.AddNewTempWasteFromCollector(r.Context(), p); err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }
	}

	fmt.Printf("\n%s\n\n", wasteID.String())

	e := fmt.Sprintf("glass: %.2f, paper :%.2f,plastics:%.2f,metal:%.2f, funnyName: %s, uuid:%s", p.GlassWeighed, p.PaperWeighed, p.PlasticsWeighed, p.MetalWeighed, p.FunnyName, wasteID.String())
	fmt.Println(e)
	qrCode := NewSimpleQRCode(e)
	codeData, err := qrCode.generateWithWatermark()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := agentPages.RenderQRCode(codeData).Render(context.Background(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func getFormattedData(r *http.Request) (glass, paper, metal, plastics float32, err error) {

	paperStringData := r.PostForm.Get("paper")
	if paperStringData == "" {
		paperStringData = "0"
	}
	glassStringData := r.PostForm.Get("glass")
	if glassStringData == "" {
		glassStringData = "0"
	}
	metalStringData := r.PostForm.Get("metal")
	if metalStringData == "" {
		metalStringData = "0"
	}
	plasticsStringData := r.PostForm.Get("plastics")
	if plasticsStringData == "" {
		plasticsStringData = "0"
	}

	var (
		paperFloat64, glassFloat64, metalFloat64, plasticsFloat64 float64
	)

	if paperFloat64, err = strconv.ParseFloat(paperStringData, 32); err != nil {
		return 0, 0, 0, 0, errors.New("invalid paper value")
	}
	if glassFloat64, err = strconv.ParseFloat(glassStringData, 32); err != nil {
		return 0, 0, 0, 0, errors.New("invalid glass value")
	}
	if metalFloat64, err = strconv.ParseFloat(metalStringData, 32); err != nil {
		return 0, 0, 0, 0, errors.New("invalid metal value")
	}
	if plasticsFloat64, err = strconv.ParseFloat(plasticsStringData, 32); err != nil {
		return 0, 0, 0, 0, errors.New("invalid plastics value")
	}

	return float32(glassFloat64), float32(paperFloat64), float32(metalFloat64), float32(plasticsFloat64), nil
}

type simpleQRCode struct {
	content   string
	size      int
	watermark []byte
}

const DefaultQRCodeSize = 256

func NewSimpleQRCode(content string) *simpleQRCode {
	return &simpleQRCode{
		content:   content,
		size:      DefaultQRCodeSize,
		watermark: beltechWatermarkLogo,
	}
}

func (code *simpleQRCode) generate() ([]byte, error) {
	qrCode, err := qrcode.Encode(code.content, qrcode.Medium, code.size)
	if err != nil {
		return nil, fmt.Errorf("could not generate a QR code: %v", err)
	}
	return qrCode, nil
}

func (code *simpleQRCode) generateWithWatermark() ([]byte, error) {
	qrCode, err := code.generate()
	if err != nil {
		return nil, err
	}

	qrCode, err = code.addWatermark(qrCode, code.watermark)
	if err != nil {
		return nil, fmt.Errorf("could not add watermark to QR code: %v", err)
	}

	return qrCode, nil
}

func (code *simpleQRCode) addWatermark(qrCode []byte, watermarkData []byte) ([]byte, error) {
	qrCodeData, err := png.Decode(bytes.NewBuffer(qrCode))
	if err != nil {
		return nil, fmt.Errorf("could not decode QR code: %v", err)
	}

	watermarkWidth := uint(float64(qrCodeData.Bounds().Dx()) * 0.25)
	watermark, err := resizeWatermark(bytes.NewBuffer(watermarkData), watermarkWidth)
	if err != nil {
		return nil, fmt.Errorf("could not resize the watermark image: %v", err)
	}

	watermarkImage, err := png.Decode(bytes.NewBuffer(watermark))
	if err != nil {
		return nil, fmt.Errorf("could not decode watermark: %v", err)
	}

	var halfQrCodeWidth, halfWatermarkWidth int = qrCodeData.Bounds().Dx() / 2, watermarkImage.Bounds().Dx() / 2
	offset := image.Pt(
		halfQrCodeWidth-halfWatermarkWidth,
		halfQrCodeWidth-halfWatermarkWidth,
	)

	watermarkImageBounds := qrCodeData.Bounds()
	m := image.NewRGBA(watermarkImageBounds)

	draw.Draw(m, watermarkImageBounds, qrCodeData, image.Point{}, draw.Src)
	draw.Draw(
		m,
		watermarkImage.Bounds().Add(offset),
		watermarkImage,
		image.Point{},
		draw.Over,
	)

	watermarkedQRCode := bytes.NewBuffer(nil)
	png.Encode(watermarkedQRCode, m)

	return watermarkedQRCode.Bytes(), nil
}

func resizeWatermark(watermark io.Reader, width uint) ([]byte, error) {
	decodedImage, err := png.Decode(watermark)
	if err != nil {
		return nil, fmt.Errorf("could not decode watermark image: %v", err)
	}

	m := resize.Resize(width, 0, decodedImage, resize.Lanczos3)
	resized := bytes.NewBuffer(nil)
	png.Encode(resized, m)

	return resized.Bytes(), nil
}
