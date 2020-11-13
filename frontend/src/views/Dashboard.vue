<template>
  <div>
    <AppHeader/>
    <Sidebar/>
    <router-view/>
    <!-- <AppFooter/> -->
  </div>
</template>

<script lang="ts">

import { Component, Vue } from 'vue-property-decorator'
import AppHeader from '@/components/AppHeader.vue'
import Sidebar from '@/components/Sidebar.vue'
import AppFooter from '@/components/AppFooter.vue'

@Component({
  components: {
    AppHeader,
    Sidebar,
    AppFooter
  }
}
)

export default class Dashboard extends Vue {

}
</script>

<style scoped>

</style>
