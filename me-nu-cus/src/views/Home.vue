<template>
  <v-app>
    <Header :order="order" :t_no="t_no" :ordertotal="ordertotal" :billtotal="billtotal" :bill="bill" :getdate="getdate"
      :dateid="dateid" @empty="clearOrder" @delete="deleteMenu" @incrementOrder="incrementOrder($event)"
      @minusOrder="minusOrder($event)" @createOrder="createOrder" @createBill="createBill" />
    <v-main>
      <v-container>
        <Promotion :ads="ads" />

        <v-row class="d-flex justify-end ">
          <v-col cols="auto">

            <v-chip class="ma-2" large text-color="white" @click="rankStatus()" color="orange">
              <v-avatar left>
                <v-icon>fas fa-crown</v-icon>
              </v-avatar>
              Top 10 Food
            </v-chip>
          </v-col>

          <v-col cols="5">
            <v-text-field label="Search" placeholder="Type here.." class="mt-1" outlined type="text" v-model="search">
            </v-text-field>
          </v-col>

        </v-row>

        <!-- <Menu :menu="menu" @order="addOrder" /> -->
        <Menu :menu="menu" :searchmenu="searchmenu" :rank_menu="rank_menu" :rank_status="rank_status" @order="addOrder">
        </Menu>
        <!-- <div></div> -->
        <Footer />

      </v-container>
    </v-main>
  </v-app>
</template>

<script>
  import axios from 'axios'
  import Promotion from '@/components/Promotion'
  import Menu from '@/components/Menu copy'
  import Header from '@/components/HeaderView'
  import Footer from '@/components/Footer'


  export default {
    components: {
      Menu,
      Promotion,
      Header,
      Footer

    },
    // sockets:{
    //   connect(){
    //     console.log('socket connect')
    //   }
    // },
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
        rank_menu: [],
        rank_status: false,
        ordertotal: 0,
        billtotal: 0,
        bill_ido: "",
        dateid: "",
        search: "",
        current_bill_id: "",
        first: true,
        connection: null,
        ads: [],
        key: "",
        t_no: ""
      }
    },
    created() {
      this.t_no = this.$route.params.tID
      this.getkey()
    },

    mounted() {

      // this.id= this.$route.params.id

      // axios
      //   .get(process.env.VUE_APP_BASE_URL+'/food/cus/menu-food')
      //   .then(res => {
      //     this.menu = res.data.data.item.food
      //     this.rank_menu = res.data.data.item.ranking
      //     this.ads = res.data.data.item.promotion
      //     console.log(res.data)
      //   })

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
      addOrder(name, price) {
        this.dateid = this.getdate
        if (this.first) {
          this.order.push({
            food_name: name,
            price: price,
            quantity: 1,
            is_done: false,
            bill_id: this.dateid
          });
          this.ordertotal += (price * 1);

        } else {
          this.order.push({
            food_name: name,
            price: price,
            quantity: 1,
            // bill_id: this.getdate
            bill_id: this.current_bill_id
          });
          this.ordertotal += (price * 1);
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
        axios.post(process.env.VUE_APP_BASE_URL + '/bill/create-order', fd)
          .then(res => {
            console.log(res)
          }).catch(error => {
            console.log(error)
          })
        this.bill = this.bill.concat(this.order)
        this.billtotal += this.ordertotal
        this.order = []
        this.ordertotal = 0

      },

      createBill() {
        this.bill_id = this.dateid
        var t_no = parseInt(this.t_no)
        const bill = {
          "id": this.bill_id,
          "is_paid": false,
          "done": false,
          "tableNO": t_no,
          "amount": this.ordertotal,
        };
        console.log("Bill" + bill)
        const jsonbill = JSON.stringify(bill)
        console.log("Bill JSON " + jsonbill)
        if (this.first) {
          axios.post(process.env.VUE_APP_BASE_URL + '/bill/create-bill', jsonbill)
            .then(res => {
              console.log(res)
            }).catch(error => {
              console.log(error)
            })
          this.current_bill_id = this.bill_id
          this.first = false
        }

      },

      rankStatus() {
        this.rank_status = !this.rank_status;
      },

      getdata() {
        axios
          .get(process.env.VUE_APP_BASE_URL + '/food/cus/menu-food')
          .then(res => {
            this.menu = res.data.data.item.food
            this.rank_menu = res.data.data.item.ranking
            this.ads = res.data.data.item.promotion
            console.log(res.data)
          })
      },

      getkey() {
        axios.get(process.env.VUE_APP_BASE_URL + '/table/get-table')
          .then(res => {
            for (let i = 0; i < res.data.data.item.table.length; i++) {
              if (res.data.data.item.table[i].table_no == parseInt(this.$route.params.tID)) {
                this.key = res.data.data.item.table[i].t_key
              }
            }
            if (this.$route.params.id == this.key) {
              this.getdata()
            } else {
              return this.$router.push('/404')
            }
          })
      }
    },
  }
</script>