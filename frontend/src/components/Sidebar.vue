<template>
  <div class="sidebar-page">
    <section class="sidebar-layout">
      <b-sidebar
        position="relative"
        :reduce="reduce"
        type="is-light"
        open
      >
        <div class="p-1">
          <b-menu-list label="Menu">
            <b-menu-item
              tag="router-link"
              v-for="data in sidebarData"
              :key="data.id"
              :label="$t(data.label)"
              :icon="data.icon"
              :icon_alt="data.icon_alt"
              :to="data.targetUrl"
            ></b-menu-item>
          </b-menu-list>
          <b-menu-list>
            <b-menu-item class="bottom" @click="toggle" icon="swap-horizontal" label="Toggle"></b-menu-item>
          </b-menu-list>
        </div>
      </b-sidebar>
    </section>
  </div>
</template>

<script lang="ts">
import { sidebar } from '@/json/sidebar.json'
import { Component, Prop, Vue } from 'vue-property-decorator'

@Component({})
export default class Sidebar extends Vue {
  public sidebar = JSON;
  public sidebarData = sidebar;

  private toggle (): void {
    this.reduce = !this.reduce
  }

  public expandOnHover = false;
  public mobile = 'reduce';
  public reduce = false;
}
</script>

<style lang="scss">
.bottom {
  position: absolute;
  bottom: 0;
}
.p-1 {
  padding: 1em;
  float: left;
  height: 800px;
  text-align: left;
}

// .sidebar-page {
//     display: flex;
//     flex-direction: column;
//     width: 100%;
//     min-height: 100%;
//     .sidebar-layout {
//         display: flex;
//         flex-direction: row;
//         min-height: 100%;
//     }
// }
@media screen and (max-width: 1023px) {
  .b-sidebar {
    .sidebar-content {
      &.is-mini-mobile {
        &:not(.is-mini-expand),
        &.is-mini-expand:not(:hover) {
          .menu-list {
            li {
              a {
                span:nth-child(2) {
                  display: none;
                }
              }
              ul {
                padding-left: 0;
                li {
                  a {
                    display: inline-block;
                  }
                }
              }
            }
          }
          .menu-label:not(:last-child) {
            margin-bottom: 0;
          }
        }
      }
    }
  }
}
@media screen and (min-width: 1024px) {
  .b-sidebar {
    .sidebar-content {
      &.is-mini {
        &:not(.is-mini-expand),
        &.is-mini-expand:not(:hover) {
          .menu-list {
            li {
              a {
                span:nth-child(2) {
                  display: none;
                }
              }
              ul {
                padding-left: 0;
                li {
                  a {
                    display: inline-block;
                  }
                }
              }
            }
          }
          .menu-label:not(:last-child) {
            margin-bottom: 0;
          }
        }
      }
    }
  }
}
</style>
