import Vue from 'vue'
import Vuex from 'vuex'
import {IObjectType} from "@/types/ObjectType";
import {IResourceType} from "@/types/ResourceType";
import {IUserType} from "@/types/UserType";
import {IWeatherType} from "@/types/WeatherType";
import {ISensorDataType} from "@/types/SensorDataType";
import {IChartDataType} from "@/types/ChartDataType";
import createPersistedState from 'vuex-persistedstate'


Vue.use(Vuex)

export default new Vuex.Store({
  plugins: [createPersistedState({
    storage: window.sessionStorage,
  })],
  state: {
    objects: [{
      uoi: '',
      displayName: '',
      description: '',
      geoCoordinates: {
        latitude: 0,
        longitude: 0
      },
      streetName: '',
      streetNumber: '',
      zipcode: '',
      city: '',
      country: ''
    }] as IObjectType[],

    resources: [{
      uoi: '',
      resourceId: '',
      displayName: '',
      createdAt: 0,
      resourceType: '',
      meta: null
    }] as IResourceType[],

    sensorData: [{
      device_id: '',
      timestamp: 0,
      altitude: 0,
      co2_conc: 0,
      humidity: '',
      illuminance: 0,
      motion_count: 0,
      pressure: 0,
      sound_level: 0,
      temperature: 0,
      voc_conc: 0,
    }] as ISensorDataType[],

    chartData: {
      timestamp: [],
      altitude: [],
      co2_conc: [],
      humidity: [],
      illuminance: [],
      motion_count: [],
      pressure: [],
      sound_level: [],
      voc_conc: [],
      temperature: []
    } as IChartDataType,

    object: {
      uoi: '',
      displayName: '',
      description: '',
      geoCoordinates: {
        latitude: 0,
        longitude: 0
      },
      streetName: '',
      streetNumber: '',
      zipcode: '',
      city: '',
      country: ''
    } as IObjectType,

    resource: {
      uoi: '',
      resourceId: '',
      displayName: '',
      createdAt: 0,
      resourceType: '',
      meta: null,
    } as IResourceType,
    user: {
      userId: '',
      firstname: '',
      lastname: '',
      email: '',
      birthdate: 0,
      streetName: '',
      streetNumber: '',
      zipcode: '',
      city: '',
      country: ''
    } as IUserType,

    weather: {
      temp : 0,
      humidity : 0,
      pressure : 0,
      description : '',
      rain: 0,
      weathercode : 0
    } as IWeatherType,

    cachedSensorID: 0
  },
  mutations: {},
  actions: {},
  modules: {}
})
