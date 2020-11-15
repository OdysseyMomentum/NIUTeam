<template >
    <div>

      <div class="card">
  <div class="card-content">
    <div class="media">
      <div class="media-left">
            <b-icon size="is-medium"      v-if="this.$store.state.weather.weathercode >= 200 && this.$store.state.weather.weathercode < 300" icon="weather-lightning" aria-hidden="true"/>
            <b-icon size="is-medium" v-else-if="this.$store.state.weather.weathercode >= 300 && this.$store.state.weather.weathercode < 400" icon="weather-rainy" aria-hidden="true"/>
            <b-icon size="is-medium" v-else-if="this.$store.state.weather.weathercode >= 500 && this.$store.state.weather.weathercode < 600" icon="weather-pouring" aria-hidden="true"/>
            <b-icon size="is-medium" v-else-if="this.$store.state.weather.weathercode >= 600 && this.$store.state.weather.weathercode < 700" icon="weather-snowy" aria-hidden="true"/>
            <b-icon size="is-medium" v-else-if="this.$store.state.weather.weathercode >= 700 && this.$store.state.weather.weathercode < 800" icon="weather-fog" aria-hidden="true"/>
            <b-icon size="is-medium" v-else-if="this.$store.state.weather.weathercode === 800" icon="weather-sunny" aria-hidden="true"/>
            <b-icon size="is-medium" v-else-if="this.$store.state.weather.weathercode > 800" icon="weather-partly-cloudy" aria-hidden="true"/>
      </div>
      <div class="media-content">
        <p class="title is-5">{{this.$store.state.object.city}}</p>
        <p class="subtitle is-6">{{this.$store.state.weather.temp}} °C</p>
        <p class="subtitle is-6">{{this.$store.state.weather.humidity}} % Humidity</p>
        <p class="subtitle is-6">{{this.$store.state.weather.pressure}} hPa</p>
      </div>
    </div>
    </div>
  </div>
</div>

</template>

<script lang="ts">

import { Component, Prop, Vue } from 'vue-property-decorator'
import weather from 'openweather-apis'
import { IWeatherType } from '@/types/WeatherType'

@Component
export default class Weather extends Vue {
  private cWeather: IWeatherType

  constructor () {
    super()
    weather.setLang(this.$i18n.locale)
    weather.setCity(this.$store.state.object.city)
    weather.setUnits('metric')
    // API Code
    weather.setAPPID('e69112f3b031f5aff360f016e024f9b8')
    weather.getSmartJSON((err: any, data: any) => {
      if (err) {
        console.log('an error happened')
      }
      this.$store.state.weather = data
    })
    this.cWeather = this.$store.state.weather
  }
}

</script>
