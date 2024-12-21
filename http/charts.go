package http

import (
	"bytes"
	"fmt"
	"html/template"
	"io"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/render"
)

const baseTpl = `
<div class="container">
    <div class="item" id="{{ .ChartID }}" style="height:18.5rem;"></div>
</div>
{{- range .JSAssets.Values }}
   <script src="{{ . }}"></script>
{{- end }}
<script type="text/javascript">
    "use strict";
    let goecharts_{{ .ChartID | safeJS }} = echarts.init(document.getElementById('{{ .ChartID | safeJS }}'), "{{ .Theme }}");
    let option_{{ .ChartID | safeJS }} = {{ .JSON }};
    goecharts_{{ .ChartID | safeJS }}.setOption(option_{{ .ChartID | safeJS }});
    {{- range .JSFunctions.Fns }}
    {{ . | safeJS }}
    {{- end }}

	window.addEventListener('resize', function() {
		goecharts_{{ .ChartID | safeJS }}.resize();
	});

	goecharts_{{ .ChartID | safeJS }}.setOption(%s);
	
</script>`

const barChartOptions = `{
	tooltip: {
		backgroundColor: 'rgba(50,50,50,0.01)',
		borderColor: '#000000',
		textStyle: {
			color: '#000000'
		},
		valueFormatter: function(value) {
			return value + 'Kg';
			}
		},
		xAxis: {
			name: 'waste',
			axisLabel: {
				fontSize: 10,
				fontWeight: 'bold',
			},
			nameTextStyle: {
				color: '#000000',
				fontWeight: 'bolder',
				fontSize: 10,
			}
		},
		yAxis: {
			type: 'value',
			axisLabel: {
				fontSize: 10,
				align: 'middle',
				fontWeight: '700',
				formatter: '{value} Kg'
			}
		}
	}`

type BarChartData struct {
	XAxis []string
	YAxis []float64
	Keys  []string
	Month string
	Year  string
}

func barItems(YAxis []float64, xAxis []string) []opts.BarData {
	items := make([]opts.BarData, 0, 4)
	for i := 0; i < len(xAxis); i++ {
		items = append(items, opts.BarData{
			Name:  xAxis[i],
			Value: YAxis[i],
			ItemStyle: &opts.ItemStyle{
				Color:       barChartColors[i],
				BorderColor: barChartColorsBackGround[i],
			},
		})
	}
	return items
}

var barChartColors [7]string = [7]string{
	"rgb(255, 99, 132)",
	"rgb(255, 159, 64)",
	"rgb(255, 205, 86)",
	"rgb(75, 192, 192)",
	"rgb(54, 162, 235)",
	"rgb(153, 102, 255)",
	"rgb(201, 203, 207)"}

var barChartColorsBackGround [7]string = [7]string{
	"rgba(255, 99, 132, 0.5)",
	"rgba(255, 159, 64, 0.5)",
	"rgba(255, 205, 86, 0.5)",
	"rgba(75, 192, 192, 0.5)",
	"rgba(54, 162, 235, 0.5)",
	"rgba(153, 102, 255, 0.5)",
	"rgba(201, 203, 207, 0.5)",
}

const chartName = "beltech"

func chart() template.HTML {
	return WeekStatBarChart(
		BarChartData{
			XAxis: []string{"Plastic", "Glass", "Metal", "Paper"},
			YAxis: []float64{3, 4, 6, 9},
			Month: "December",
			Year:  "2024",
			Keys:  []string{},
		},
	)
}

func WeekStatBarChart(data BarChartData) template.HTML {
	bar := charts.NewBar()

	bar.SetGlobalOptions(charts.WithInitializationOpts(
		opts.Initialization{AssetsHost: "/public/assets/scripts/"},
	))

	bar.Renderer = newchartRenderer(
		bar,
		chartName,
		fmt.Sprintf(baseTpl, barChartOptions),
		bar.Validate,
	)

	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: fmt.Sprintf("Beltech colletion for the month of %s, %s.", data.Month, data.Year),
			TitleStyle: &opts.TextStyle{
				Color:      "black",
				FontWeight: "700",
				FontSize:   10,
				FontFamily: "Poppins",
			},
			// Subtitle: fmt.Sprintf("from %s - %s %s. Total Uptime of %s", data.XAxis[0], data.XAxis[6], data.Month, utils.UsageTimeInHrsMin(data.TotalUptime)),
			// SubtitleStyle: &opts.TextStyle{
			// 	Color:      "black",
			// 	FontWeight: "500",
			// 	FontSize:   14,
			// 	FontFamily: "system-ui",
			// },
			Left: "center",
		}),
		charts.WithLegendOpts(opts.Legend{
			Show: opts.Bool(true),
			Left: "20%",
			Data: data.Keys,
			Icon: "triangle",
		}),

		// charts.WithToolboxOpts(opts.Toolbox{
		// 	Feature: &opts.ToolBoxFeature{
		// 		SaveAsImage: &opts.ToolBoxFeatureSaveAsImage{
		// 			Show:  opts.Bool(true),
		// 			Title: "Download chart as Image",
		// 			Name:  fmt.Sprintf("weekly-screentime-%s-%s-%s-%s", data.XAxis[0], data.XAxis[3], data.Month, data.Year),
		// 		},
		// 	},
		// }),
		// charts.WithTooltipOpts(opts.Tooltip{
		// 	Show:      opts.Bool(true),
		// 	Trigger:   "axis",
		// 	TriggerOn: "mousemove",
		// 	AxisPointer: &opts.AxisPointer{
		// 		Type: "cross",
		// 	},
		// }),
	)
	bar.SetXAxis(data.XAxis).
		AddSeries("Waste", barItems(data.YAxis, data.XAxis)).SetSeriesOptions()
	return renderToHtml(bar)
}

type Renderer interface {
	Render(io.Writer) error
}

func renderToHtml(c render.Renderer) template.HTML {
	var buf bytes.Buffer
	_ = c.Render(&buf)
	return template.HTML(buf.String())
}

type chartRenderer struct {
	c             interface{}
	before        []func()
	chartName     string
	chartTemplate string
}

func newchartRenderer(c interface{}, chartName string, chartTemplate string, before ...func()) render.Renderer {
	return &chartRenderer{
		c:             c,
		chartName:     chartName,
		chartTemplate: chartTemplate,
		before:        before}
}

func (r *chartRenderer) Render(w io.Writer) error {
	for _, fn := range r.before {
		fn()
	}

	tpl := template.
		Must(template.New(r.chartName).
			Funcs(template.FuncMap{
				"safeJS": func(s interface{}) template.JS {
					return template.JS(fmt.Sprint(s))
				},
			}).
			Parse(r.chartTemplate),
		)

	err := tpl.ExecuteTemplate(w, r.chartName, r.c)
	return err
}

// To satisfy the interface
func (r *chartRenderer) RenderContent() []byte {
	return []byte{}
	// To satisfy the interface
}
func (r *chartRenderer) RenderSnippet() render.ChartSnippet {
	return render.ChartSnippet{}
}
