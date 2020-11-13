<template>
    <section>
      <div class="box">
         <div class="locale-changer">
           {{ $t('language') }}:
    <select v-model="$i18n.locale">
      <option v-for="(lang, i) in langs" :key="`Lang${i}`" :value="lang">{{ lang }}</option>
    </select>
  </div>
  </div>
    </section>

</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'
export default {
  data () {
    return {
      langs: ['en', 'de'],
      darkMode: false
    }
  }

}
</script>

<style>
.box {
  text-align: left;
  display: block;
  margin: 1em;
  width: 350px;
  height: 800px;
  background-color: #E2E2E2;
  scrollbar-width: thin;
  overflow-y : auto;
  overflow-x: hidden;
}
.info {
  margin: 1em;
  float: left;
  width: 800px;
  height: 800px;
}

</style>
