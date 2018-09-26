<template>
  <div>
    <div id="chart-container"></div>
  </div>
</template>

<script type="javascript" src="../assets/js/dark-unica.js"></script>

<script>
  var $ = require('jquery')
  window.jQuery = $
  var Highcharts = require('highcharts/highstock')
  // var moment = require('moment')

  export default {
    name: 'chart',
    props: ['simpleDataFromTrade', 'priceFromTrade'],
    data () {
      return {
        simpleData: []
      }
    },
    methods: {
      drawChart: function () {
        /* eslint-disable */
        (function(a){"object"===typeof module&&module.exports?module.exports=a:a(Highcharts)})(function(a){a.createElement("link",{href:"https://fonts.googleapis.com/css?family\x3dUnica+One",rel:"stylesheet",type:"text/css"},null,document.getElementsByTagName("head")[0]);a.theme={colors:"#2b908f #90ee7e #f45b5b #7798BF #aaeeee #ff0066 #eeaaee #55BF3B #DF5353 #7798BF #aaeeee".split(" "),chart:{backgroundColor:null,style:{fontFamily:"'Unica One', sans-serif"},
        plotBorderColor:"#606063"},title:{style:{color:"#E0E0E3",textTransform:"uppercase",fontSize:"20px"}},subtitle:{style:{color:"#E0E0E3",textTransform:"uppercase"}},xAxis:{gridLineColor:"#707073",labels:{style:{color:"#E0E0E3"}},lineColor:"#707073",minorGridLineColor:"#505053",tickColor:"#707073",title:{style:{color:"#A0A0A3"}}},yAxis:{gridLineColor:"#707073",labels:{style:{color:"#E0E0E3"}},lineColor:"#707073",minorGridLineColor:"#505053",tickColor:"#707073",tickWidth:1,title:{style:{color:"#A0A0A3"}}},
        tooltip:{backgroundColor:"rgba(0, 0, 0, 0.85)",style:{color:"#F0F0F0"}},plotOptions:{series:{dataLabels:{color:"#B0B0B3"},marker:{lineColor:"#333"}},boxplot:{fillColor:"#505053"},candlestick:{lineColor:"white"},errorbar:{color:"white"}},legend:{itemStyle:{color:"#E0E0E3"},itemHoverStyle:{color:"#FFF"},itemHiddenStyle:{color:"#606063"}},credits:{style:{color:"#666"}},labels:{style:{color:"#707073"}},drilldown:{activeAxisLabelStyle:{color:"#F0F0F3"},activeDataLabelStyle:{color:"#F0F0F3"}},navigation:{buttonOptions:{symbolStroke:"#DDDDDD",
        theme:{fill:"#505053"}}},rangeSelector:{buttonTheme:{fill:"#505053",stroke:"#000000",style:{color:"#CCC"},states:{hover:{fill:"#707073",stroke:"#000000",style:{color:"white"}},select:{fill:"#000003",stroke:"#000000",style:{color:"white"}}}},inputBoxBorderColor:"#505053",inputStyle:{backgroundColor:"#333",color:"silver"},labelStyle:{color:"silver"}},navigator:{handles:{backgroundColor:"#666",borderColor:"#AAA"},outlineColor:"#CCC",maskFill:"rgba(255,255,255,0.1)",series:{color:"#7798BF",lineColor:"#A6C7ED"},
        xAxis:{gridLineColor:"#505053"}},scrollbar:{barBackgroundColor:"#808083",barBorderColor:"#808083",buttonArrowColor:"#CCC",buttonBackgroundColor:"#606063",buttonBorderColor:"#606063",rifleColor:"#FFF",trackBackgroundColor:"#404043",trackBorderColor:"#404043"},legendBackgroundColor:"rgba(0, 0, 0, 0.5)",background2:"#505053",dataLabelsColor:"#B0B0B3",textColor:"#C0C0C0",contrastTextColor:"#F0F0F3",maskColor:"rgba(255,255,255,0.3)"};a.setOptions(a.theme)});
        /* eslint-enable */

        let self = this

        jQuery.noConflict()

        Highcharts.setOptions({
          global: {
            useUTC: false
          }
        })
        // Create the chart
        var chart = Highcharts.stockChart('chart-container', {
          chart: {
            height: 500,
            events: {
              load: function () {
                // set up the updating of the chart each second
                var series = this.series[0]
                setInterval(function () {
                  var x = (new Date()).getTime() // current time
                  var y = self.priceFromTrade
                  series.addPoint([x, y], true, true)
                }, 60000)
              }
            }
          },

          title: {
            text: 'BTC/ETH'
          },

          rangeSelector: {
            buttons: [{
              count: 10,
              type: 'minute',
              text: '10M'
            }, {
              count: 60,
              type: 'minute',
              text: '1H'
            }, {
              count: 1,
              type: 'day',
              text: '1D'
            }, {
              type: 'all',
              text: '1W'
            }],
            inputEnabled: false,
            selected: 3
          },

          series: [{
            name: 'BTC/USD',
            data: self.simpleData,
            type: 'area',
            threshold: null,
            tooltip: {
              valueDecimals: 2
            }
          }],

          responsive: {
            rules: [{
              condition: {
                maxWidth: 500
              },
              chartOptions: {
                chart: {
                  height: 300
                },
                subtitle: {
                  text: null
                },
                navigator: {
                  enabled: false
                }
              }
            }]
          }
        })

        $('#small').click(function () {
          chart.setSize(400)
        })

        $('#large').click(function () {
          chart.setSize(800)
        })

        $('#auto').click(function () {
          chart.setSize(null)
        })
      }
    },
    mounted () {
      let self = this

      let decide = setInterval(function () {
        if (self.simpleDataFromTrade.length > 10) {
          let simpleData = []
          let dDate

          self.simpleDataFromTrade.forEach(function (d) {
            let simpleDataItem = []

            dDate = d[0] * 1000

            simpleDataItem[0] = dDate
            simpleDataItem[1] = d[1]

            simpleData.push(simpleDataItem)
          })
          self.simpleData = simpleData

          self.drawChart()
          self.$emit('chartLoaded', false)
          clearInterval(decide)
        } else {
          console.log('waiting for chart data')
        }
      }, 2000)
    }
  }
</script>

<style scoped>
</style>
