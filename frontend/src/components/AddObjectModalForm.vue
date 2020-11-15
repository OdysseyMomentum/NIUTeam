<template>
    <div class="modal-card" style="width: auto">
      <header class="modal-card-head">
        <p class="modal-card-title">Add Object</p>
        <button type="button" class="delete" @click="$emit('close')" />
      </header>
      <section class="modal-card-body">
          <fieldset>
              <div class="columns is-mobile is-multiline">
                  <div class="column is-full pt-0"><label>Name</label><input class="input" v-model="displayName"></div>
                  <div class="column is-full pt-0"><label>Description</label><input class="input" v-model="description"></div>
                  <div class="column is-full pt-0"><label>UOI</label><input class="input" placeholder="UOI..." v-model="uoi"></div>
        <div class="column is-four-fifths pt-0">
                <label>Street Name</label
                ><input
                  class="input"
                  placeholder="Street Name..."
                  v-model="streetName"
                />
              </div>
              <div class="column is-one-fifths pt-0">
                <label>Number</label
                ><input
                  class="input"
                  placeholder="Number..."
                  v-model="streetNumber"
                />
              </div>
              <div class="column is-four-fifths pt-0 pt-0">
                <label>City</label
                ><input
                  name="city"
                  class="input"
                  placeholder="City..."
                  v-model="city"
                />
              </div>
              <div class="column is-one-fifths pt-0">
                <label>Zipcode</label
                ><input
                  name="zipcode"
                  class="input"
                  placeholder="Zipcode..."
                  v-model="zipcode"
                />
              </div>
              <div class="column is-full pt-0">
                <label>Country</label
                ><input
                  name="country"
                  class="input"
                  placeholder="Country..."
                  v-model="country"
                />
              </div>
              </div>
          </fieldset>
          </section>
      <footer class="modal-card-foot" style="justify-content: flex-end;">
        <button @click="addObject(), $emit('close'), toast()" class="button is-primary is-pulled-right">Add</button>
      </footer>
    </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { uuid } from 'vue-uuid'
export default class Modalform extends Vue {
  private displayName = ''
  private description = ''
  private uoi = uuid.v4()
  private streetName = ''
  private streetNumber = ''
  private city = ''
  private zipcode = ''
  private country = ''

  toast () {
    this.$buefy.toast.open({
      message: 'Ressource added, please reload the page!',
      type: 'is-success'
    })
  }

  async addObject () {
    const accessToken = await this.$auth.getTokenSilently({})
    const result = await fetch(encodeURI('/api/object/add'), {
      method: 'POST',
      headers: {
        Authorization: 'Bearer ' + accessToken
      },
      body: JSON.stringify({
        uoi: this.uoi,
        description: this.description,
        displayName: this.displayName,
        streetName: this.streetName,
        streetNumber: this.streetNumber,
        city: this.city,
        zipcode: this.zipcode,
        country: this.country
      })
    })
  }
}
</script>
