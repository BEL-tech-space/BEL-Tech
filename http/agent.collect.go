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
	"time"

	agentPages "beltech/frontend/pages/agent"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/nfnt/resize"
	"github.com/skip2/go-qrcode"
)

const qrCodeContentFormat = "glass:%.2f#metal:%.2f#plastics:%.2f#paper:%.2f#txnname:%s#uuid:%s#points:%.2f"

func (a *App) collectPage(w http.ResponseWriter, r *http.Request) {
	agentID := r.Context().Value(contextKeyAgentID).(uuid.UUID)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := agentPages.CollectPage(agentID.String()).Render(context.Background(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) generateQRCode(w http.ResponseWriter, r *http.Request) {

	agentID := chi.URLParam(r, "agentID")
	if agentID == "" {
		a.clientError(w, http.StatusBadRequest, errors.New("agentID not specified"))
		return
	}

	agentUUID, err := uuid.Parse(agentID)
	if err != nil {
		a.clientError(w, http.StatusBadRequest, errors.New("invalid agentID"))
		return
	}

	agentId := r.Context().Value(contextKeyAgentID).(uuid.UUID)
	if agentUUID.String() != agentId.String() {
		a.clientError(w, http.StatusBadRequest, errors.New("collector is not who was expected"))
		return
	}

	collectorEmail := r.Context().Value(contextKeyEmail).(string)

	if err = r.ParseForm(); err != nil {
		a.clientError(w, http.StatusBadRequest, err)
		return
	}

	glass, paper, metal, plastics, err := getFormattedData(r)
	if err != nil {
		a.clientError(w, http.StatusBadRequest, err)
		return
	}

	if sum := glass + paper + metal + plastics; sum <= 0 {
		a.clientError(w, http.StatusBadRequest, errors.New("zero values"))
		return
	}

	wasteLog := service.WasteLog{
		WasteID:         uuid.New(),
		GlassWeighed:    glass,
		MetalWeighed:    metal,
		PlasticsWeighed: plastics,
		PaperWeighed:    paper,
		TxnName:         gofakeit.Name(),
		AgentID:         agentUUID,
		CollectorEmail:  collectorEmail,
	}

	if err = includePointsEarned(&wasteLog); err != nil {
		a.serverError(w, fmt.Errorf("error calculating points earned: error: %w ", err))
		return
	}

	if saved := a.wasteLogCache.SetWithTTL(
		wasteLog.WasteID.String(),
		wasteLog,
		1,
		24*time.Hour,
	); !saved {
		a.serverError(w, errors.New("waste data failed saving to cache"))
		return
	}

	content := fmt.Sprintf(qrCodeContentFormat,
		wasteLog.GlassWeighed,
		wasteLog.MetalWeighed,
		wasteLog.PlasticsWeighed,
		wasteLog.PaperWeighed,
		wasteLog.TxnName,
		wasteLog.WasteID.String(),
		wasteLog.TotalPointsEarned,
	)

	qrCode := NewSimpleQRCode(content)
	codeData, err := qrCode.generateWithWatermark()
	if err != nil {
		a.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := agentPages.RenderQRCode(codeData).Render(context.Background(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func includePointsEarned(wasteLog *service.WasteLog) error {
	wasteLog.GlassPointsEarned = wasteLog.GlassWeighed * 1.2
	wasteLog.MetalPointsEarned = wasteLog.MetalWeighed * 0.2
	wasteLog.PaperPointsEarned = wasteLog.PaperWeighed * 0.9
	wasteLog.PlasticsPointsEarned = wasteLog.PlasticsWeighed * 1.5
	wasteLog.TotalPointsEarned = float64(wasteLog.GlassPointsEarned) + float64(wasteLog.MetalPointsEarned) +
		float64(wasteLog.PaperPointsEarned) + float64(wasteLog.PlasticsPointsEarned)

	return nil
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
