<template>
  <div class="hello" :style="{ backgroundImage: '@/assets/background.jpeg' }">
    <div class="overlay">
      <div class="columns">
<!--        First Column-->
        <div class="column is-6">
        </div>
        <!--        Third Column-->
        <div class="column to-center is-5">
          <div class="box shadowed">
            <article class="media">
              <div class="media-left">
                <figure class="image is-64x64">
                  <img src="../assets/filedgr-ico.png">
                </figure>
              </div>
              <div class="media-content">
                <div class="content">
                  <p class="aligned-left">
                    <strong>Filedgr</strong>
                    <br>
                    At Filedgr, pronounced <span style="color: #46bed7; font-weight: bold;">[faɪl ˈlɛʤər]</span> file ledger, we believe in Security, Transparency, and Integrity. This is why we build our solution, enabling the exchange of data with proof of authenticity and delivery. Where existing blockchain solutions fail at providing true decentralization, and therewith true transparency, for information other than financial transactions, Filedgr fills the void.
                  </p>
                </div>
              </div>
            </article>
          </div>
        </div>
        <div class="column" >
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'

@Component
export default class HelloWorld extends Vue {
  created () {
    console.log(window.innerHeight)
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}

.box {
  margin-left: 50px;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: inline-block;
  margin: 0 10px;
}

a {
  color: #42b983;
}

p.aligned-left{
  text-align: left;
}

.columns {
  margin-left: 0px;
  margin-right: 0px;
}

.to-center{
  margin-top: 200px;
}

.shadowed {
  box-shadow: 2px 2px 4px #000000;
}

.hello {
  background-image: url('../assets/background.jpeg');
  width: 100%;
  height: 84vh;
  background-repeat:no-repeat;
  background-position: center center;
}

.overlay {
  z-index: 1;
  height: 100%;
  width: 100%;
  /* position: fixed; */
  overflow: hidden;
  top: 0px;
  left: 0px;
  background: linear-gradient(to bottom, rgba(255, 255, 255, 0),#40c2d8);
}
</style>
