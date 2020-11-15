<template>
  <div>
    <AppHeader/>
    <Sidebar/>
    <router-view/>
  </div>
</template>

<script lang="ts">

import { Component, Vue } from 'vue-property-decorator'
import AppHeader from '@/components/AppHeader.vue'
import Sidebar from '@/components/Sidebar.vue'

@Component({
  components: {
    AppHeader,
    Sidebar
  }
}
)

export default class Dashboard extends Vue {
}
</script>

<style scoped>

</style>
