<template>
  <div class="card info">
    <Map/>
    <div class="card-content">
      <p class="title">
        {{ this.object.displayName }}
      </p>
      <b-tabs type="is-boxed">
        <b-tab-item label="Object" icon="home">
          <fieldset>
            <div class="columns is-mobile is-multiline">
              <div class="column is-full pt-0">
                <label><strong>UOI</strong></label
                ><input class="input" readonly v-model="this.object.uoi"/>
              </div>
              <div class="column is-full pt-0">
                <label><strong>Description</strong></label
                ><input readonly class="input" v-model="this.object.description"/>
              </div>
              <div class="column is-four-fifths pt-0">
                <label><strong>Street Name</strong></label
                ><input
                name="street"
                class="input"
                readonly
                v-model="this.object.streetName"
              />
              </div>
              <div class="column is-one-fifths pt-0">
                <label><strong>Number</strong></label
                ><input class="input" readonly v-model="this.object.streetNumber"/>
              </div>
              <div class="column is-four-fifths pt-0 pt-0">
                <label><strong>City</strong></label
                ><input class="input" readonly v-model="this.object.city"/>
              </div>
              <div class="column is-one-fifths pt-0">
                <label><strong>Zipcode</strong></label
                ><input class="input" readonly v-model="this.object.zipcode"/>
              </div>
              <div class="column is-full pt-0">
                <label><strong>Country</strong></label
                ><input class="input" readonly v-model="this.object.country"/>
              </div>
            </div>
          </fieldset>
        </b-tab-item>
        <b-tab-item label="Sensors" icon="leak">
          <div class="columns is-multiline">
            <div class="column is-one-fifth field">
              <b-button class="  is-ghost is-fullwidth mb-2" v-for="resources in this.getSensors()"
                        :key="resources.resourceId" :label="resources.displayName"
                        @click="getSensorData(resources.resourceId, startTime, endTime)"></b-button>
            </div>
            <div class="column is-four-fifths">
              <div class="columns is-multiline">
                <div class="column is-one-third">
                  <b-field custom-class="is-small" horizontal label="Start:">
                    <b-datetimepicker v-model="startTime"
                                      size=is-small
                                      horizontal-time-picker='true'
                                      :max-datetime="maxDatetime"
                                      placeholder="Start Date...">
                    </b-datetimepicker>
                  </b-field>
                </div>
                <div class="column is-one-third">
                  <b-field custom-class="is-small" horizontal label="End:">
                    <b-datetimepicker v-model="endTime"
                                      size=is-small
                                      horizontal-time-picker='true'
                                      :max-datetime="maxDatetime"
                                      placeholder="End Date...">
                    </b-datetimepicker>
                  </b-field>
                </div>
                <div class="column is-2">
                  <b-button @click="getSensorData($store.state.cachedSensorID, startTime, endTime )"
                            class="is-small is-fullwidth">Refresh
                  </b-button>
                </div>
                <div class="column">
                  <Chart class="ml-6" :key="rerenderKey" v-if="loaded"/>
                </div>
              </div>
            </div>
          </div>

        </b-tab-item>
        <b-tab-item label="Documents" icon="file-document-outline">
          <a class="panel-block" v-for="resources in this.getDocuments()"
             :key="resources.resourceId"
             :href="resources.meta.access"
             @click="startDownload( resources.resourceId )">
            <b-icon v-if="resources.meta.filetype === '.pdf'" icon="file-pdf-box" aria-hidden="true"/>
            <b-icon v-else-if="resources.meta.filetype === '.docx' || resources.meta.filetype === '.doc'"
                    icon="file-word-box" aria-hidden="true"/>
            <b-icon v-else-if="resources.meta.filetype === '.xlsx' || resources.meta.filetype === '.xlsx'"
                    icon="file-excel-box" aria-hidden="true"/>
            <b-icon v-else-if="resources.meta.filetype === '.pptx' || resources.meta.filetype === '.ppt'"
                    icon="file-powerpoint-box" aria-hidden="true"/>
            <b-icon v-else-if="resources.meta.filetype === '.png' || resources.meta.filetype === '.jpg'"
                    icon="file-image" aria-hidden="true"/>
            {{ resources.meta.filename }} {{ resources.meta.filesize | prettyBytes(2, false, 'MB') }}
          </a>
        </b-tab-item>
        <b-tab-item label="Weather" icon="weather-partly-cloudy">
          <Weather/>
        </b-tab-item>
        <b-tab-item label="Add User" icon="account-plus">
          <div class="column is-one-third">
          <b-field>
            <template slot="label">
              Email
              <b-tooltip type="is-white" label="Add a User to see your Object">
                <b-icon size="is-small" icon="help-circle-outline"></b-icon>
              </b-tooltip>
            </template>
            <b-input placeholder="Email..." v-model="newUserEmail" ></b-input>
          </b-field>
            <b-button @click="addUser()" type="is-success is-pulled-right mt-4"
            >+
            </b-button
            >
          </div>
        </b-tab-item>
      </b-tabs>
    </div>
  </div>
</template>

<script lang="ts">
import 'leaflet/dist/leaflet.css'
import Chart from '@/components/plots/Chart.vue'
import Map from '@/components/Map.vue'
import Weather from '@/components/Weather.vue'
import { Component, Vue } from 'vue-property-decorator'
import { IDocumentMeta, IResourceType } from '@/types/ResourceType'
import { IObjectType } from '@/types/ObjectType'
import { ISensorDataType } from '@/types/SensorDataType'

@Component({
  components: {
    Map,
    Weather,
    Chart
  }
})
export default class Details extends Vue {
  private newUserEmail = ''
  private objectResources: IResourceType[]
  private object: IObjectType
  private rerenderKey = 0
  public loaded = false

  private endTime = new Date()
  private startTime = new Date(Date.now() - 43200 * 1000)
  private maxDatetime = new Date()

  constructor () {
    super()
    this.objectResources = this.$store.state.resources
    this.object = this.$store.state.object
  }

  private successUserAdd () {
    this.$buefy.toast.open({
      message: 'User ' + this.newUserEmail + ' added!',
      type: 'is-success',
      position: 'is-bottom'
    })
  }

  private async addUser () {
    const accessToken = await this.$auth.getTokenSilently({})
    fetch(encodeURI('/api/user/addToObject'), {
      method: 'POST',
      headers: {
        Authorization: 'Bearer ' + accessToken
      },
      body: JSON.stringify({
        object: {
          uoi: this.object.uoi
        },
        user: {
          email: this.newUserEmail
        }
      })
    }).then(async response => {
      const data = await response.json() as JSON
      console.log(data)
      this.successUserAdd()
    })
      .catch(error => {
        console.error('There was an error!', error)
      })
  }

  private async getSensorData (resourceId: number, startTime: Date, endTime: Date) {
    this.$store.state.cachedSensorID = resourceId
    const accessToken = await this.$auth.getTokenSilently({})
    fetch(encodeURI('/api/telemetry/get'), {
      method: 'POST',
      headers: {
        Authorization: 'Bearer ' + accessToken
      },
      body: JSON.stringify({
        object: {
          uoi: this.object.uoi
        },
        resource: {
          resourceId: resourceId
        },
        range: {
          begin: Math.round(startTime.getTime() / 1000),
          end: Math.round(endTime.getTime() / 1000)
        }

      })
    }).then(async response => {
      const sensorData: ISensorDataType[] = await response.json() as ISensorDataType[]

      this.$store.state.chartData.timestamp = sensorData.map(i => i.timestamp * 1000)
      this.$store.state.chartData.altitude = sensorData.map(i => i.altitude)
      this.$store.state.chartData.co2_conc = sensorData.map(i => i.co2_conc)
      this.$store.state.chartData.humidity = sensorData.map(i => i.humidity)
      this.$store.state.chartData.illuminance = sensorData.map(i => i.illuminance)
      this.$store.state.chartData.motion_count = sensorData.map(i => i.motion_count)
      this.$store.state.chartData.pressure = sensorData.map(i => i.pressure)
      this.$store.state.chartData.sound_level = sensorData.map(i => i.sound_level)
      this.$store.state.chartData.temperature = sensorData.map(i => i.temperature)
      this.$store.state.chartData.voc_conc = sensorData.map(i => i.voc_conc)
      this.loaded = true
      this.rerenderKey += 1
    })
      .catch(error => {
        console.error('There was an error!', error)
      })
  }

  private async startDownload (resourceId: string) {
    console.log('Downloading: ' + resourceId)
    const accessToken = await this.$auth.getTokenSilently({})
    fetch(encodeURI('/api/resource/get'), {
      method: 'POST',
      headers: {
        Authorization: 'Bearer ' + accessToken
      },
      body: JSON.stringify({
        object: {
          uoi: this.object.uoi
        },
        resource: {
          resourceId: resourceId
        }

      })
    }).then(async response => {
      const resource: IResourceType = await response.json() as IResourceType
      const resourceMeta = resource.meta as IDocumentMeta

      if (resourceMeta.access !== undefined) {
        if (resourceMeta.filetype === '.jpg' ||
          resourceMeta.filetype === '.jpeg' ||
          resourceMeta.filetype === '.png' ||
          resourceMeta.filetype === '.pdf' ||
          resourceMeta.filetype === '.svg') {
          window.open(resourceMeta.access, '_blank')
        } else {
          window.open(resourceMeta.access)
        }
      }
    })
      .catch(error => {
        console.error('There was an error!', error)
      })
  }

  private getDocuments (): IResourceType[] {
    return this.objectResources.filter(r => (r.resourceType === 'DOCUMENT' && r.meta != null))
  }

  private getSensors (): IResourceType[] {
    return this.objectResources.filter(r => (r.resourceType === 'SENSOR'))
  }
}
</script>

<style lang="scss" scoped>
.card-content {
  text-align: left;
}
</style>
