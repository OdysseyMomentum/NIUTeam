<template>
  <div>
    <l-map
      :zoom="zoom"
      :center="center"
      :options="mapOptions"
      @update:center="centerUpdate"
      @update:zoom="zoomUpdate"
      style="z-index: 0; height: 250px;"
    >
      <l-tile-layer
        :url="url"
        :attribution="attribution"
      />
      <l-marker :lat-lng="coordinates"
                :icon="icon">
      </l-marker>
    </l-map>
  </div>
</template>

<script>
import { latLng, icon } from 'leaflet'
import { LMap, LTileLayer, LMarker } from 'vue2-leaflet'
import 'leaflet/dist/leaflet.css'

export default {
  name: 'Map',
  components: {
    LMap,
    LTileLayer,
    LMarker
  },

  data () {
    return {
      zoom: 16,
      center: latLng(this.$store.state.object.geoCoordinates.latitude, this.$store.state.object.geoCoordinates.longitude),
      url: 'https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png',
      attribution:
        '&copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors',
      coordinates: latLng(this.$store.state.object.geoCoordinates.latitude, this.$store.state.object.geoCoordinates.longitude),
      currentZoom: 11.5,
      currentCenter: latLng(this.$store.state.object.geoCoordinates.latitude, this.$store.state.object.geoCoordinates.longitude),
      showParagraph: false,
      mapOptions: {
        zoomSnap: 0.5
      },
      icon: icon({
        iconUrl: require('../assets/marker.svg'), // marker.svg is the normal one
        iconSize: [50, 100],
        iconAnchor: [16, 37]
      })
    }
  },
  methods: {
    zoomUpdate (zoom) {
      this.currentZoom = zoom
    },
    centerUpdate (center) {
      this.currentCenter = center
    }
  }

}
</script>
