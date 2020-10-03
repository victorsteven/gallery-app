<template>
<div>
    <div class="page-container md-layout-column">
     <md-toolbar class="md-primary">
        <md-button style="color: #039be5">Gallery</md-button>
     </md-toolbar>
    <md-content class="main-body">
       <div>
          <slot name="body"></slot>
        </div>
    </md-content>
    
    <div id="footer">
      <Footer/>
    </div>
    </div>
  </div>
</template>

<script>
    import Footer from "./Footer";

    export default {
        name: "Layout",
        components: {Footer},
         data: () => ({
            showNavigation: false,
            showSidepanel: false
        }),
        computed: {
          authenticated(){
            return this.$store.state.authenticated
          }
        },
        methods: {
          logout(){
            this.$store.dispatch('logout').then(() => {
                this.$router.push("/")
            })
          }
        }
    }
</script>

<style scoped>

a {
  text-decoration: none;
}
a:hover {
  text-decoration: none;
}
a:link {
  text-decoration: none;
}

a:active {
  text-decoration: none;
}
.logout-style {
  cursor: pointer;
  color: #039be5;
}
#footer {
   position:absolute;
   bottom:0!important;
   width:100%;
   height:60px;   /* Height of the footer */
   background:#6cf;
   margin-bottom: 0px!important;
}
.md-elevation-4 {
  box-shadow: none!important
}

.main-body {
  background-image: 
    linear-gradient(to bottom, rgba(255, 255, 255, 0.5)),
    url("https://res.cloudinary.com/chikodi/image/upload/v1523699895/back.jpg");
}
</style>
