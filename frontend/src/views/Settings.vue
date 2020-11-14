<template>
  <section>
    <div class="tile is-ancestor">
      <div class="tile is-parent is-vertical">
        <article class="tile is-child box is-2">
          <div class="content">
            <div class="locale-changer">
              {{ $t("language") }}:
              <select v-model="$i18n.locale">
                <option
                  v-for="(lang, i) in langs"
                  :key="`Lang${i}`"
                  :value="lang"
                >
                  {{ lang }}
                </option>
              </select>
            </div>
          </div>
        </article>
      </div>
    </div>
  </section>
</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";
export default {

  data() {
    return {
      langs: ["en", "de"],
      darkMode: false,
    };
  },
};
</script>

<style scoped>
</style>
