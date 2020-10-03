<template>
    <div class="card">
        <div class="card-image">
            <img :src="img.imgUrl"  class="image-style">
            <span v-if="imageInfo">
                <div class="untag">
                    <button @click="unTag(img.id)" class="button button-red">UnTag</button>
                </div>
            </span>
            <span v-else>
                <div class="tag">
                    <button @click="tag(img.id)" class="button button-green">Tag</button>
                </div>
            </span>
        </div>
    </div>
</template>


<script>

import axios from 'axios'
import API_ROUTE from '../../.env'

export default {

    props: {
        userIp: {
            type: String,
            required: true,
        },
        img: {
            type: Object,
            required: true
        },
    },
     
    data: () => ({
        imageInfo: '',
    }),

    mounted() {
        this.findTag(this.img.id)
    },


    methods: {

        async tag(imageId){

            const payload = {
                imageId: imageId,
                userIp: this.userIp
            }

            try {

                await axios.post(`${API_ROUTE}/image_info`, payload)

                this.imageInfo = true

            } catch(err) {

                console.log("this is the error tagging: ", err)

            }
        },


        async unTag(imageId){

            try {

                await axios.delete(`${API_ROUTE}/image_info/${imageId}`)

                this.imageInfo = false
        

            } catch(err) {

                console.log("the error untagging: ", err)

            }
        },

        async findTag(imageId){

            try {

                const res = await axios.get(`${API_ROUTE}/image_info/${imageId}`)

                this.imageInfo = res.data.body
        
            } catch(err) {

                console.log("error getting the tag: ", err)

            }
        },
    }
}

</script>

<style  scoped>

 .image-style {
     height:220px
   }

   .tag {
       padding: 5px;
   }
   .tag > a {
       color: green;

   }

   .untag {
       padding: 5px;
   }
   .untag > a {

       color: red;
   }

   .button {
        background-color: #4CAF50; 
        border: none;
        color: white;
        padding: 8px 10px;
        text-align: center;
        text-decoration: none;
        display: inline-block;
        margin: 2px 2px;
        transition-duration: 0.4s;
        cursor: pointer;
    }

    .button-green {
        background-color: white; 
        color: black; 
        border: 2px solid #4CAF50;
    }

    .button-green:hover {
        background-color: #4CAF50;
        color: white;
    }

    .button-red {
        background-color: white; 
        color: black; 
        border: 2px solid #f44336;
    }

    .button-red:hover {
        background-color: #f44336;
        color: white;
    }

</style>