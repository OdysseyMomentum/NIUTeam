<template>
  <section>
    <div class="tile is-ancestor">
      <div class="tile is-4 is-vertical is-parent">
        <div class="tile is-child box has-text-left">
          <div class="media">
            <figure class="image is-32x32">
              <img src="../assets/user.svg"/>
            </figure>
            <p class="title is-size-5 mt-2 ml-2">Personal Details</p>
          </div>
          <fieldset>
            <div class="columns is-mobile is-multiline">
              <div class="column is-half pt-0 mt-4">
                <label>First Name</label
                ><input
                class="input"
                placeholder="First Name..."
                v-model="cUser.firstname"
              />
              </div>
              <div class="column is-half pt-0 mt-4">
                <label>Last Name</label
                ><input
                name="lname"
                class="input"
                placeholder="Last Name..."
                v-model="cUser.lastname"
              />
              </div>
              <div class="column is-full pt-0">
                <label>Email</label
                ><input
                name="email"
                type="email"
                class="input"
                placeholder="Email Address..."
                v-model="cUser.email"
                disabled
              />
              </div>
              <div class="column is-full pt-0">
                <label>Birthday</label
                ><input
                name="bday"
                type="date"
                class="input"
                placeholder="Click to select..."
                v-model="formatedBirthdate"
              />
              </div>
              <div class="column is-three-quarters pt-0">
                <label>Street Name</label
                ><input
                name="street"
                class="input"
                placeholder="Street Name..."
                v-model="cUser.streetName"
              />
              </div>
              <div class="column is-one-quarters pt-0">
                <label>Number</label
                ><input
                name="snumber"
                class="input"
                placeholder="Number..."
                v-model="cUser.streetNumber"
              />
              </div>
              <div class="column is-three-quarters pt-0 pt-0">
                <label>City</label
                ><input
                name="city"
                class="input"
                placeholder="City..."
                v-model="cUser.city"
              />
              </div>
              <div class="column is-one-quarters pt-0">
                <label>Zipcode</label
                ><input
                name="postal_code"
                class="input"
                placeholder="Zipcode..."
                v-model="cUser.zipcode"
              />
              </div>
              <div class="column is-full pt-0">
                <label>Country</label
                ><input
                name="country"
                class="input"
                placeholder="Country..."
                v-model="cUser.country"
              />
              </div>
            </div>
          </fieldset>
          <b-button @click="putUser()" type="is-success is-pulled-right mt-4"
          >Update
          </b-button
          >
        </div>
      </div>
    </div>
  </section>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import { IUserType } from '@/types/UserType'
import dateFormat from "dateformat";

@Component
export default class Profile extends Vue {
  private cUser: IUserType
  private formatedBirthdate: string
  
  constructor () {
    super()
    this.cUser = this.$store.state.user
    this.formatedBirthdate = dateFormat(new Date(this.cUser.birthdate * 1000), 'yyyy-mm-dd')
  }

  async putUser () {
    const accessToken = await this.$auth.getTokenSilently()
    console.log(accessToken)
    const result = await fetch(encodeURI('/api/user/update'), {
      method: 'PUT',
      headers: {
        Authorization: 'Bearer ' + accessToken
      },
      body: JSON.stringify({
        userId: this.$auth.user.sub,
        email: this.$auth.user.email,
        firstname: this.cUser.firstname,
        lastname: this.cUser.lastname,
        birthdate: parseInt((new Date(this.formatedBirthdate).getTime() / 1000).toFixed(0)),
        streetName: this.cUser.streetName,
        streetNumber: this.cUser.streetNumber,
        zipcode: this.cUser.zipcode,
        city: this.cUser.city,
        country: this.cUser.country
      })
    })
    const data = await result.json()
    console.log(data)
  }
}
</script>
