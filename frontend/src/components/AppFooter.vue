<template>
  <div>
    <footer class="footer">
      <div class="content has-text-centered">
        <p>
          <strong>Filedgr</strong> is a product of <a href="https://niugroup.lu">NIUGroup SARLS</a>.
        </p>
      </div>
    </footer>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'

@Component
export default class AppFooter extends Vue {}
</script>

<style scoped>
.footer {
  background-color: #ffffff;
  padding: 30px 30px 30px 30px;
  height: 50px;
}
</style>
