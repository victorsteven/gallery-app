<template>
  <Layout>
    <div slot="body" class="cont">
      <div class="preloader-background" v-if="loading">
        <p class="blinking" style="font-size: 40px; color: #304ffe">Loading...</p>
      </div>
      <div class="row pad-images" v-if="images.length > 0">
        <div class="col s12 m3 style-render" v-for="(img, index) in images" :key="index">
          <OneImage :userIp="userIp" :img="img" />
        </div>
      </div>
    </div>
  </Layout>
</template>



<script>

import Layout from '../Nav/Layout'
import OneImage from './OneImage'
import API_ROUTE from '../../.env'
import axios from 'axios'

export default {

    components: { Layout, OneImage },

    data: () => ({
        loading: true,
        userIp: "",
        images: [],
    }),

    mounted() {

      let userIp = localStorage.getItem("user_ip")
      if (userIp !== null) {
        this.userIp = userIp
          this.getImages()
      } else {
        this.getUserIp()
      }
    },

    computed: {
        disabled(){
            return this.loading === true
        },
    },

    methods: {

      async getUserIp() {

        try {

          this.loading = true

          const res = await axios.get('https://api.ipify.org?format=json&callback=?');

          this.userIp = res.data["ip"]
          localStorage.setItem("user_ip", this.userIp)

          this.getImages()

        } catch(err) {

          console.log("this is the error getting the user ip: ", err)

          this.loading = false

        }
    },

    async getImages() {

      try {

        this.loading = true

        const res = await axios.get(`${API_ROUTE}/images`)

        this.loading = false

          this.images =  res.data.images

      } catch(err) {

        console.log("this is the error getting the images: ", err)

        this.loading = false

      }
    },
  }
}
</script>

<style scoped>
   .card__meta i.small {
       font-size: 1.5rem;
   }
   .card .card-content .card-title {
       color: #0D8080 !important;
       line-height: 18px;
   }
   .card .card-title {
       font-size: 18px;
       font-weight: 400;
   }
   .card .card-content p{
     color: #1E1E1E;
     font-size: 13px;
   }
   .style-render {
      padding: 20px 10px 0px 10px;
   }

  @media only screen and (min-width: 768px) {
    .pad-images {
      padding:  10px 10px;
    }
  }


  .cont {
      padding-bottom: 50px;
  }
  .style-render[data-v-364c61fe] {
    padding: 5px 10px 0px 10px;
}
  .preloader-background {
    display: flex;
    align-items: center;
    justify-content: center;
    background-color: #fff;
    position: fixed;
    z-index: 999;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
  }
  .preloader-background.p {
    padding-top:120px;
    margin-left: -60px;
    opacity: 0.8;
  } 
  .blinking {
      animation: blinker 1.5s linear infinite;
  }
  @keyframes blinker {  
    50% { opacity: 0; }
  }
</style>