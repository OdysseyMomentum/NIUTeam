<template>
  <div>
    <AppHeader/>
    <HelloWorld/>
    <AppFooter/>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import HelloWorld from '@/components/HelloWorld.vue'
import AppHeader from '@/components/AppHeader.vue'
import AppFooter from '@/components/AppFooter.vue'

@Component({
  components: {
    HelloWorld,
    AppHeader,
    AppFooter
  }
})

export default class Home extends Vue {

}
</script>
