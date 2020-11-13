<template>
<div class="chart">
  <figure id="lineChart"></figure>
  <button @click="reDraw">Click me!</button>
</div>
</template>
<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'
import * as d3 from 'd3'

@Component({})
export default class LineChart extends Vue {
  private data = [
    { price: 40129, year: 2009 },
    { price: 166443, year: 2010 },
    { price: 150793, year: 2012 },
    { price: 62342, year: 2014 },
    { price: 27647, year: 2016 },
    { price: 30000, year: 2017 }
  ];

  private svg: any;
  private d?: number;
  private margin = 50;
  private width = 750 - (this.margin * 2);
  private height = 400 - (this.margin * 2);

  private createSvg (): void {
    this.svg = d3.select('#lineChart')
      .append('svg')
      .attr('width', this.width + (this.margin * 2))
      .attr('height', this.height + (this.margin * 2))
      .append('g')
      .attr('transform', 'translate(' + this.margin + ',' + this.margin + ')')
  }

  private drawPlot (inputData: any[]): void {
    const price = inputData.map(item => parseInt(item.price))
    const year = inputData.map(item => item.year)
    const priceMin = Math.min(...price)
    const priceMax = Math.max(...price)

    // Add X axis
    const x = d3.scaleLinear()
      .domain([2009, 2017])
      .range([0, this.width])
    this.svg.append('g')
      .attr('transform', 'translate(0,' + this.height + ')')
      .call(d3.axisBottom(x).tickFormat(d3.format('d')))

    // Add Y axis
    const y = d3.scaleLinear()
      .domain([0, priceMax])
      .range([this.height, 0])
    this.svg.append('g')
      .call(d3.axisLeft(y))

    // Add line
    this.svg.append('path')
      .datum(inputData)
      .attr('fill', 'none')
      .attr('stroke', 'black')
      .attr('stroke-width', 2)
      .attr('d', d3.line()
        // .curve(d3.curveNatural)
        .x(function (d: any) { return x(d.year) })
        .y(function (d: any) { return y(d.price) })
      )

    // Add dots
    const dots = this.svg.append('g')
    dots.selectAll('dot')
      .data(inputData)
      .enter()
      .append('circle')
      .attr('cx', (d: { year: any}) => x(d.year))
      .attr('cy', (d: { price: any}) => y(d.price))
      .attr('r', 3)
      // .style('opacity', 0.5)

    this.svg.append('text')
      .attr('x', (this.width / 2))
      .attr('y', 0 - (this.margin / 2))
      .attr('text-anchor', 'middle')
      .text('Price')
  }

  private clearPlot () {
    d3.select('svg').remove()
  }

  private reDraw (): void {
    this.clearPlot()
    this.createSvg()
    this.drawPlot(
      [
        { price: 40000 * Math.random(), year: 2009 },
        { price: 40000 * Math.random(), year: 2010 },
        { price: 40000 * Math.random(), year: 2012 },
        { price: 40000 * Math.random(), year: 2014 },
        { price: 40000 * Math.random(), year: 2016 },
        { price: 40000 * Math.random(), year: 2017 }
      ]
    )
  }

  mounted () {
    this.createSvg()
    this.drawPlot(this.data)
  }
}
</script>

<style>
.chart {
  padding: 1em;
  float: left;
  margin: 1em;
  background-color: #f8f8f8;
}

.p {
  text-align: center;
}
</style>
