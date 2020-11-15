<template>
  <div style="height: 300px">
    <canvas id="chart"></canvas>
    <div class="buttons is-centered">
      <a @click="createChart('Temperature', $store.state.chartData.temperature, 23)" class="button is-small is-multiline">Temperature</a>
      <a @click="createChart('CO2 Concentration', $store.state.chartData.co2_conc, 1400)"
         class="button is-small is-multiline">CO2</a>
      <a @click="createChart('Humidity', $store.state.chartData.humidity)"
         class="button is-small is-multiline">Humidity</a>
      <a @click="createChart('Illuminance', $store.state.chartData.illuminance)" class="button is-small is-multiline">Illuminance</a>
      <a @click="createChart('Pressure', $store.state.chartData.pressure)"
         class="button is-small is-multiline">Pressure</a>
      <a @click="createChart('Motion', $store.state.chartData.motion_count)" class="button is-small  is-multiline">Motion</a>
      <a @click="createChart('Sound Level', $store.state.chartData.sound_level)" class="button is-small is-multiline">Sound
        level</a>
      <a @click="createChart('Altitude', $store.state.chartData.altitude)"
         class="button is-small is-multiline">Altitude</a>
      <a @click="createChart('VOC Concentration', $store.state.chartData.voc_conc)"
         class="button is-small is-multiline">VOC</a>
    </div>
  </div>

</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'
import Chart from 'chart.js'
import annotation from 'chartjs-plugin-annotation'
Chart.plugins.register(annotation)

@Component({})
export default class LineChart extends Vue {
  createChart (labelToDisplay: string, dataToDisplay: [], thresholdValue: any) {
    const ctx: any = document.getElementById('chart')
    const thresholdHighArray = new Array(dataToDisplay.length).fill(thresholdValue)
    const myChart = new Chart(ctx, {
      type: 'line',
      data: {
        labels: this.$store.state.chartData.timestamp,
        datasets: [
          {
            data: dataToDisplay,
            borderColor: '#4AD9D9',
            borderWidth: 2,
            pointRadius: 0,
            pointHitRadius: 0,
            fill: false
          }, {
            data: thresholdHighArray,
            pointRadius: 0,
            borderColor: 'rgba(63, 191, 63, 0.16)',
            backgroundColor: 'rgba(63, 191, 63, 0.1)'
          }

        ]
      },
      options: {
        tooltips: {
          enabled: false
        },
        legend: {
          display: false
        },
        title: {
          display: true,
          text: labelToDisplay
        },
        responsive: true,
        maintainAspectRatio: false,
        plugins: {
          annotation: {
            annotations: [{
              id: 'a-line-1',
              type: 'line',
              mode: 'horizontal',
              scaleID: 'y',
              value: '25',
              borderColor: 'red',
              borderWidth: 2
            }]
          }
        },
        scales: {
          xAxes: [{
            id: 'x',
            type: 'time',
            time: { displayFormats: { minute: 'hA' } },
            display: true,
            scaleLabel: {
              display: true
            },
            ticks: {
              source: 'auto'
            }
          }],
          yAxes: [{
            id: 'y',
            display: true,
            scaleLabel: {
              display: true
            }
          }]
        }
      }
    })
  }

  mounted () {
    this.createChart('Temperature', this.$store.state.chartData.temperature, 23)
  }
}
</script>
