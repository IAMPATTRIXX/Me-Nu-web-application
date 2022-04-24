<template>
  <v-app>
    <Header :order="order" :ordertotal="ordertotal" :billtotal="billtotal"  :bill="bill" :getdate="getdate" :dateid="dateid" @empty="clearOrder"
      @delete="deleteMenu" @incrementOrder="incrementOrder($event)" @minusOrder="minusOrder($event)" @createOrder="createOrder" @createBill="createBill"/>
    <v-main>
      <v-container>
        <Promotion />

        <v-row class="d-flex justify-end ">
          <v-col cols="5">
            <v-text-field label="Outlined" placeholder="Placeholder" outlined type="text" v-model="search">
            </v-text-field>
          </v-col>
        </v-row>


        <!-- <Menu :menu="menu" @order="addOrder" /> -->
        <Menu :menu="menu" :searchmenu="searchmenu" @order="addOrder"></Menu>

      </v-container>
    </v-main>
  </v-app>
</template>

<script>
  import axios from 'axios'
  import Promotion from '@/components/Promotion'
  import Menu from '@/components/Menu'
  import Header from '@/components/HeaderView'

  export default {
    components: {
      Menu,
      Promotion,
      Header
    },
    // data: () => ({
    //   menu: [],
    //   order: [],
    //   ordertotal: 0,
    //   bill_ido: "",
    //   dateid: "",
    //   search:"",
    //   return:{
    //     search:""
    //   }
    // }),

    // }),
    data() {
      return {
        menu: [],
        order: [],
        bill: [],
        ordertotal: 0,
        billtotal:0,
        bill_ido: "",
        dateid: "",
        search: "",
        current_bill_id:"",
        first:true
      }
    },


    mounted() {
      axios
        .get('http://localhost:8081/food/cus/menu-food')
        .then(res => {
          this.menu = res.data
          // console.log(this.menu)
          // console.log('Success')
          // console.log('------------------------------------------------------')
          this.menu = this.menu.data.item.food
          // console.log(this.menu)
        })
      
    },

    computed: {
      getdate() {
        const d = new Date();
        let day = d.getDate();
        let month = d.getMonth();
        let year = d.getFullYear();
        let hour = d.getHours();
        let min = d.getMinutes();
        let sec = d.getSeconds();
        let days = day.toString();
        let months = month.toString();
        let years = year.toString();
        let hours = hour.toString();
        let mins = min.toString();
        let secs = sec.toString();
        return days.concat(months, years, hours, mins, secs)
      },

      searchmenu() {
        return this.menu.filter(menus =>
          menus.name.toLowerCase().includes(this.search.toLowerCase())
        );
      }
    },

    methods: {
      addOrder(menus) {
        let checkNotExist = true;
        this.dateid = this.getdate
        // this.order.forEach(function (value) {
        //     if (value.name.toLowerCase().match(menus.name.toLowerCase())) {
        //         checkNotExist = true;
        //         value.count ++;
        //         self.ordertotal += (value.price * 1);
        //     }
        // })
        if (checkNotExist) {
          this.order.push({
            food_name: menus.name,
            price: menus.price,
            quantity: 1,
            bill_id: this.dateid,
            is_done: false
          });
          this.ordertotal += (menus.price * 1);
          // console.log(this.order)
          // สร้างไว้ก่อนน
          // ---------------------------------------------------------
          // if (first){
          //   this.order.push({
          //   food_name: menus.name,
          //   price: menus.price,
          //   quantity: 1,
          //   is_done: false
          //   bill_id: this.dateid
          // });
          // this.ordertotal += (menus.price * 1);
          // }else{
          //   this.order.push({
          //   food_name: menus.name,
          //   price: menus.price,
          //   quantity: 1,
          //   // bill_id: this.getdate
          //   bill_id: this.current_bill_id
          // });
          // this.ordertotal += (menus.price * 1);
          // }
          // ------------------------------------------------------------
        }

      },
      deleteMenu(orderindex) {
        if (orderindex || orderindex === 0) {
          let count = this.order[orderindex].quantity;
          let price = this.order[orderindex].price;
          this.order.splice(orderindex, 1);
          this.ordertotal -= (count * price);
        }
      },
      clearOrder() {
        this.order = [];
        this.ordertotal = 0;
      },
      incrementOrder(orderIndex) {
        this.order[orderIndex].quantity++;
        this.ordertotal += (this.order[orderIndex].price * 1);
      },
      minusOrder(orderIndex) {
        if (this.order[orderIndex].quantity > 1) {
          this.order[orderIndex].quantity--;
          this.ordertotal -= (this.order[orderIndex].price * 1);
        } else if (this.order[orderIndex].quantity == 1) {
            if (orderIndex || orderIndex === 0) {
              let count = this.order[orderIndex].quantity;
              let price = this.order[orderIndex].price;
              this.order.splice(orderIndex, 1);
              this.ordertotal -= (count * price);
          }
        }
      },
      
      createOrder() {
        const fd = JSON.stringify(this.order)
        
        //สร้างไว้ก่อนยังไม่ได้ลอง
        //------------------------------------------------------------
        // if(this.first){
        //   axios.post('http://localhost:8081/bill/create-order', fd)
        //   .then(res => {
        //     console.log(res)
        //   }).catch(error => {
        //     console.log(error)
        //   })
        //   
        //   this.bill = this.order.slice()
        //   console.log(this.bill)
        // }else{
        //   axios.post('http://localhost:8081/bill/create-order', fd)
        //   .then(res => {
        //     console.log(res)
        //   }).catch(error => {
        //     console.log(error)
        //   })
        //   this.bill = this.order.slice()
        //   console.log(this.bill)
        // }
        // -----------------------------------------------------------

        // const bodyFormData = new FormData();
        // this.order.forEach((item) => {
        //   bodyFormData.append('order', item);
        // });
        // console.log(bodyFormData)
        // console.log("fd" + fd)
        // console.log(this.order)
        
        //ใช้ได้จริง
        //---------------------------------------------------------
        axios.post('http://localhost:8081/bill/create-order', fd)
          .then(res => {
            console.log(res)
          }).catch(error => {
            console.log(error)
          })
          console.log("1111")
        // -----------------------------------------------------
          
          // if(this.first){
          //   this.bill = this.order.slice()
          //   this.billtotal = this.ordertotal
          //   console.log(this.bill)
          // }else{
          //   this.bill = this.bill.concat(this.order.slice)
          //   console.log("สั่งซ้ำนะ")
          // }
      },


      createBill() {
        // const d = new Date();
        // const text = d.toISOString();
        this.bill_id = this.dateid
        const bill = {
          "id": this.bill_id,
          "is_paid": false,
          "done": false,
          "table_no": 1,
          "amount": this.ordertotal,
        };
        // const test = JSON.stringify(fdd)
        // console.log("test :" + test)
        // ใช้ได้
        // ---------------------------------------------------------
        // axios.post('http://localhost:8081/bill/create-bill', bill)
        //   .then(res => {
        //     console.log(res)
        //   }).catch(error => {
        //     console.log(error)
        //   })
        // --------------------------------------------------------
          // this.first = false
          // console.log(this.first)


        //---------------------------------------------------------------
        // สร้างไว้ก่อน
        if(this.first){
          axios.post('http://localhost:8081/bill/create-bill', bill)
          .then(res => {
            console.log(res)
          }).catch(error => {
            console.log(error)
          })
         this.current_bill_id = this.bill_id
         this.first = false 
        }
        //----------------------------------------------------------------
      },



    },
  }
</script>