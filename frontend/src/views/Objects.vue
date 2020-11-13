<template>
<div>
  <div class="box">
    <section >
        <b-collapse
              :open="false"
              style="width: 300px"
              class="card mar-bot"
              :class="{highlight:data.id == id}"
              animation="slide"
              v-for="data in dummyObjectsData"
              :key="data.id"
            >
            <div
                slot="trigger"
                slot-scope="props"
                class="card-header"
                role="button"
                @click="focus(data.id)">
                <p v-if='editing === data.id' class="card-header-title">
                    <b-input type="text" v-model.lazy="data.displayName" />
                </p>
                <p v-else class="card-header-title">
                    {{data.displayName}}
                </p>
                <a class="card-header-icon">
                    <b-icon
                        :icon="props.open ? 'menu-down' : 'menu-up'">
                    </b-icon>
                </a>
             </div>
              <div
                slot="trigger"
                slot-scope="props"
                class="card-header"
                role="button"
                @click="focus(data.id)">
                <p v-if='editing === data.id' class="card-header-title">
                    <b-input type="text" v-model="data.displayName" />
                </p>
                <p v-else class="card-header-title">
                    {{data.displayName}}
                </p>
                <a class="card-header-icon">
                    <b-icon
                        :icon="props.open ? 'menu-down' : 'menu-up'">
                    </b-icon>
                </a>
             </div>
            <div class="card-content" @click="focus(data.id)">
                <div v-if='editing === data.id' class="content">
                    Description:<b-input type="text" v-model="data.description" /><br>
                    Tags:<b-taginput v-model="data.tags" ellipsis icon="label" placeholder="Add a tag">
                    <p class="content"><b></b> {{ tags }}</p>
            </b-taginput>
                </div>
                <div v-else class="content">
                    ID: {{data.id}} <br>
                    {{ $t('Creation Date') }}: {{data.creationDate}} <br>
                    {{ $t('Location') }}: {{data.location}} <br>
                    {{ $t('Description') }}: {{data.description}} <br>
                    Tags: <b-tag class="mar-left" v-for="tag in data.tags" :key="tag">{{tag}}</b-tag>
                </div>
            </div>
            <footer v-if='editing === data.id' class="card-footer">
                <a class="card-footer-item" @click="saveThing(data.id)">{{ $t('Save') }}</a>
                <a class="card-footer-item" @click="cancelEditing(data.id)">{{ $t('Cancel') }}</a>
            </footer>
                <footer v-else class="card-footer">
                <a class="card-footer-item" @click="editThing(data.id)">{{ $t('Edit') }}</a>
                <a class="card-footer-item" @click="confirmDelete(data.id)">{{ $t('Delete') }}</a>
            </footer>
        </b-collapse>
    </section>
  </div>
  <!-- right side -->
  <div class="info"
    :class="{display:hidden == true}">
  <div class="card-content">
    <p class="title center">
      {{displayName}}
    </p>
    <p class="subtitle">
      Id: {{this.id}} <br>
      {{ $t('Creation Date') }}: {{this.creationDate}}<br>
      {{ $t('Description') }}: {{this.description}}<br>
      Tags: <b-tag class="mar-left" v-for="tag in this.tags" :key="tag">{{tag}}</b-tag>
    </p>
  </div>
  <LineChart/>
</div>
</div>
</template>
<script lang="ts">
import { dummyObjects } from '../json/dummyObjects.json'
import LineChart from '@/components/plots/LineChart.vue'
import { Component, Prop, Vue } from 'vue-property-decorator'

@Component({
  components: {
    LineChart
  }
}
)
export default class Things extends Vue {
  public dummyObjects = JSON;
  public dummyObjectsData = dummyObjects;
  public id = -1;
  public editing =-1;
  public displayName = '';
  public creationDate = '';
  public description = '';
  public tags = [];
  public cachedDisplayName = '';
  public cachedDesc = '';
  public cachedTags = [];
  public hidden = true;

  public focus (id: number): void {
    this.id = this.dummyObjectsData[id].id
    this.displayName = this.dummyObjectsData[id].displayName
    this.creationDate = this.dummyObjectsData[id].creationDate
    this.description = this.dummyObjectsData[id].description
    this.tags = this.dummyObjectsData[id].tags
    this.hidden = false
  }

  public confirmDelete (id: number): void {
    this.$buefy.dialog.confirm({
      title: 'Deleting Thing',
      message: 'Are you sure you want to <b>delete</b> your Thing? This action cannot be undone.',
      confirmText: 'Delete Thing',
      type: 'is-danger',
      hasIcon: true
      // onConfirm: () => this.dummyThingsData = this.dummyThingsData.filter(
      //   dummyThingsData => dummyThingsData[id].id !== id
      // )
    })
  }

  public editThing (id: number): void {
    this.id = this.dummyObjectsData[id].id
    this.editing = this.id
    this.cachedDisplayName = this.dummyObjectsData[id].displayName
    this.cachedDesc = this.dummyObjectsData[id].description
    this.cachedTags = this.dummyObjectsData[id].tags
    console.log(this.cachedTags)
    console.log(this.dummyObjectsData[id].tags)
  }

  public cancelEditing (id: number): void {
    this.id = this.dummyObjectsData[id].id
    this.dummyObjectsData[id].displayName = this.cachedDisplayName
    this.dummyObjectsData[id].description = this.cachedDesc
    this.dummyObjectsData[id].tags = this.cachedTags
    this.editing = -1
  }

  public saveThing (id: number): void {
    this.id = this.dummyObjectsData[id].id
    this.displayName = this.dummyObjectsData[id].displayName
    this.description = this.dummyObjectsData[id].description
    this.tags = this.dummyObjectsData[id].tags
    this.editing = -1
  }
}
</script>

<style lang="scss">

.box {
  padding: auto;
  margin: 1em;
  float: left;
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

.highlight {
  background: lightyellow;
  margin-bottom: 10px !important ;
}

.display {
  display: none;
}

.center {
  text-align: center;
}

.card-content{
  text-align: left;
}

.mar-left{
  margin-left: 5px;
}

.mar-bot{
  margin-bottom: 10px;
}
</style>
