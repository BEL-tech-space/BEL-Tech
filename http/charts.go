package http

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"time"

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
	(function() {
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
		})();
	
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
			axisLabel: {
				fontSize: 12,
				color: '#232323',
				lineHeight: 17.7,
				fontFamily: 'TT Chocolates',
				align: 'middle',
				fontWeight: '500',
			}
		},
		yAxis: {
			type: 'value',
			axisLabel: {
				fontSize: 12,
				color: '#232323',
				lineHeight: 17.7,
				fontFamily: 'TT Chocolates',
				align: 'middle',
				fontWeight: '500',
				formatter: '{value} Kg',				
				margin: 14, 
			}
		},
		series: [{
			barCategoryGap: '85%',
			itemStyle: {
        		borderRadius: [10, 10, 0, 0]
      		},
		}],
		grid: {
			bottom: '12%',
			top: '12%'
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
				Color:       "rgba(1, 105, 52, 1)",
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

func produceChartHTML(values [4]float64, t time.Time) template.HTML {
	return WeekStatBarChart(
		BarChartData{
			XAxis: []string{"Plastic", "Glass", "Metal", "Paper"},
			YAxis: values[:],
			Month: t.Month().String(),
			Year:  fmt.Sprintf("%d", t.Year()),
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
			Title: fmt.Sprintf("Month of %s, %s.", data.Month, data.Year),
			TitleStyle: &opts.TextStyle{
				Color:      "black",
				FontWeight: "700",
				FontSize:   10,
				FontFamily: "TT Chocolates",
			},
			Left: "center",
		}),
		charts.WithLegendOpts(opts.Legend{
			Show: opts.Bool(true),
			Left: "20%",
			Data: data.Keys,
			Icon: "triangle",
		}),
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
