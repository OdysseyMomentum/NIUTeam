<template>
  <div>
    <div class="box">
      <section>
        <b-collapse
          :open="false"
          style="width: 340px"
          class="card mar-bot"
          animation="slide"
          v-for="object in this.$store.state.objects"
          :key="object.displayName"
        >
          <div
            slot="trigger"
            slot-scope="props"
            class="card-header"
            role="button"
            @click="getObject(object.uoi), getResources(object.uoi)"
          >
            <p class="card-header-title">
              {{ object.displayName }}
            </p>
            <a class="card-header-icon">
              <b-icon :icon="props.open ? 'menu-down' : 'menu-up'"> </b-icon>
            </a>
          </div>
          <div
            slot="trigger"
            slot-scope="props"
            class="card-header"
            role="button"
            @click="getObject(object.uoi), getResources(object.uoi)"
          >
            <p class="card-header-title">
              {{ object.displayName }}
            </p>
            <a class="card-header-icon">
              <b-icon :icon="props.open ? 'menu-down' : 'menu-up'"> </b-icon>
            </a>
          </div>
          <div class="card-content" @click="getObject(object.uoi), getResources(object.uoi)">
            <div class="content">
              UOI: <b-tag class="is-pulled-right"> {{ object.uoi }} </b-tag>
              <br />
              {{ $t("Description") }}: {{ object.description }} <br />
            </div>
          </div>
        </b-collapse>
        <div class="footer-button"><b-button
          @click="openModal()"
          style="position: absolute; right: 0; bottom: 0"
          type="is-success is-pulled-right mb-4 mr-4"
          >+</b-button
        ></div>
      </section>
    </div>
    <Details  v-if="show" :key="rerenderKey" />
  </div>
</template>
<script lang="ts">
import 'leaflet/dist/leaflet.css'
import Details from '@/components/Details.vue'
import { Component, Vue } from 'vue-property-decorator'
import { IResourceType } from '@/types/ResourceType'
import { IObjectType } from '@/types/ObjectType'
import AddObjectModalForm from '@/components/AddObjectModalForm.vue'

@Component({
  components: {
    Details
  }
})
export default class Things extends Vue {
  public rerenderKey = 0
  public show = false
  public isModalActive = false
  public uoi = ''
  public displayName = ''
  public description = ''

  public streetName = ''
  public city = ''

  async getObject (uoi: string) {
    const accessToken = await this.$auth.getTokenSilently({})
    const result = await fetch(encodeURI('/api/object/get'), {
      method: 'POST',
      headers: {
        Authorization: 'Bearer ' + accessToken
      },
      body: JSON.stringify({ uoi: uoi })
    })
    const response: IObjectType = await result.json() as IObjectType
    this.$store.state.object = response

    this.show = true
    // force update the map
    this.rerenderKey += 1
  }

  async getResources (uoi: string) {
    const accessToken = await this.$auth.getTokenSilently({})
    const result = await fetch(encodeURI('/api/resource/list'), {
      method: 'POST',
      headers: {
        Authorization: 'Bearer ' + accessToken
      },
      body: JSON.stringify({ uoi: uoi })
    })
    // const response: MyResourceType[] = await result.json() as MyResourceType[]
    const response: IResourceType[] = await result.json() as IResourceType[]

    this.$store.state.resources = response
  }

  openModal () {
    this.$buefy.modal.open({
      parent: this,
      component: AddObjectModalForm,
      hasModalCard: true,
      customClass: 'custom-class custom-class-2',
      trapFocus: true
    })
  }
}
</script>

<style lang="scss" scoped>

.footer-button {
}

.box {
  position: relative;
  width: 390px;
  padding: auto;
  margin: 1em;
  float: left;
  height: 800px;
  scrollbar-width: thin;
  overflow-y: auto;
  overflow-x: hidden;
}

.info {
  margin: 1em;
  float: left;
  width: 60vw;
  height: 800px;
}

.display {
  display: none;
}

.center {
  text-align: center;
}

.card-content {
  text-align: left;
}

.mar-left {
  margin-left: 5px;
}

.mar-bot {
  margin-bottom: 10px;
}
</style>
