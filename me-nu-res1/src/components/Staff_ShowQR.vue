<template>
  <div>
   <!-- <v-card>
    <v-img src="https://www.thesprucepets.com/thmb/sfuyyLvyUx636_Oq3Fw5_mt-PIc=/3760x2820/smart/filters:no_upscale()/adorable-white-pomeranian-puppy-spitz-921029690-5c8be25d46e0fb000172effe.jpg" width="480" height="320">
    </v-img>
   </v-card> -->
   
    <v-img  v-show="this.status == true" :src="code" width="300" height="300">
    </v-img>

  

   
    <!-- <span> {{test}}</span> -->
  </div>
</template>

<script>
  import QRCode from 'qrcode'
  import axios from 'axios'

  export default {
    data: () => ({
      code: "",
      link: "",
      tID: 0,
      status:false,
      data: []


    }),

    mounted() {
      
    },

    created() {
      // แบบอัตโนมัติ ปล.โค้ดโง่
      // setInterval(() => {
      //   this.getdata();
      // },30000000)

      this.getdata();
      



    },
    methods: {
      getdata() {
        this.tID = this.$route.query.tID
        axios.get(process.env.VUE_APP_BASE_URL+'/table/get-table')
          .then(res => {
            console.log(res.data.data.item)
            this.data = res.data.data.item.table[this.tID - 1]
            this.link = this.data.url
            this.status = this.data.status
            if(this.data.status == true){
              console.log('ได้รับ status = true แล้ว')
              // var opts ={
              //   small : true,
              //   size
              // }
              let linkgen = this.link
              QRCode.toDataURL(linkgen,{width : 200})
                .then(url => {
                  this.code = url
                })
            }
          }).catch(error => {
            console.log(error)
          })
      },
    }
  }
</script>