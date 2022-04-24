<template>
  <div>
    <v-app-bar color="blue-grey lighten-2" dense dark>
      

      <v-toolbar-title>Me-nU : {{t_no}}</v-toolbar-title>

      <v-spacer></v-spacer>

      <v-dialog v-model="billinfo" width="500" :retain-focus="false">
        <template v-slot:activator="{ on, attrs }">
          <v-btn v-bind="attrs" v-on="on" class="ma-2" color="primary" dark>
            Bill
            <v-icon dark right> mdi-food</v-icon>
          </v-btn>
        </template>
        <v-card>
          <v-card-title class="text-h5 grey lighten-2">
            Your Bill
          </v-card-title>

          <v-divider></v-divider>

          <v-container>
            <v-simple-table>
              <template v-slot:default>
                <thead>
                  <tr>
                    <th></th>
                    <th class="text-left">Name</th>
                    <th class="text-left">Price</th>
                    <th class="text-left">Qty</th>
                  </tr>
                </thead>
                <tbody>
                  <tr v-for="(item, index) in bill" :key="index">
                    <td>
                      <!-- <v-hover color="red">
                        <v-btn color="red accent-1" elevation="2" dark small class="white--text px-0 py-0"
                          min-width="35" height="35" @click="$emit('delete', index)">
                          <v-icon>mdi-delete</v-icon>
                        </v-btn>
                      </v-hover> -->
                    </td>
                    <td>{{ item.food_name }}</td>
                    <td>${{ item.price }}</td>
                    <td>{{ item.quantity}}</td>

                    <td>
                      <v-dialog v-model="reviewinfo" width="500">
                        <template v-slot:activator="{ on, attrs }">
                          <v-btn icon v-on="on" v-bind="attrs" @click="getcurrent(item.food_name,item.bill_id)">
                            <v-icon>mdi-pencil</v-icon>
                          </v-btn>

                        </template>

                        <v-card>
                          <!-- <v-card-title class="text-h5 grey lighten-2">
                            Review here...
                          </v-card-title> -->
                          <v-divider></v-divider>

                          <v-card-text>
                            <v-textarea name="review" auto-grow clearable clear-icon="mdi-close-circle"
                              prepend-icon="fas fa-pencil" v-model="review">
                            </v-textarea>
                          </v-card-text>

                          <v-divider></v-divider>

                          <v-card-actions>
                            <v-spacer></v-spacer>
                            <v-btn color="green" icon text @click="onCreatedReview">
                              <v-icon>mdi-checkbox-marked-circle</v-icon>
                              <!-- <font-awesome-icon icon="fa-solid fa-circle-check" /> -->
                            </v-btn>
                            <v-btn color="red" icon @click="reviewinfo = false">
                              <v-icon>mdi-close-circle</v-icon>
                            </v-btn>
                          </v-card-actions>
                        </v-card>
                      </v-dialog>

                    </td>

                  </tr>
                </tbody>
              </template>
            </v-simple-table>
            <v-row dense class="mt-5 mb-5 mr-5 d-flex flex-row-reverse">
              ${{billtotal}}
            </v-row>
          </v-container>



          <v-divider></v-divider>

          <v-card-actions>
            <v-spacer></v-spacer>
            <!-- <v-btn color="primary" text @click="getdate">
              Test
            </v-btn> -->
            <!-- <v-btn color="primary" text @click="createBill();createOrder()">
              Order
            </v-btn> -->
            <!-- <v-btn color="red" text @click="$emit('empty')">
              Clear
            </v-btn> -->
            <v-btn color="primary" text @click="billinfo = false">
              Close
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
      <!-- ปุ่มเปิด dialog ของ Order -->
      <v-dialog v-model="menuinfo" width="1000" :retain-focus="false">
        <template v-slot:activator="{ on, attrs }">
          <v-btn v-bind="attrs" v-on="on" class="ma-2" color="primary" dark>
            Order
            <v-icon dark right> mdi-food-fork-drink </v-icon>
          </v-btn>
        </template>
        <!-- dialog ของ Order -->
        <v-card>
          <v-card-title class="text-h5 grey lighten-2">
            Your Order
          </v-card-title>

          <v-divider></v-divider>
          <!-- ส่วนที่แสดงผล -->
          <!-- <v-container>
              <v-row dense v-for="(item, index) in order" :key="index">
                <v-col cols="4">
                  Name : {{item.name}}
                </v-col>
                <v-col cols="4">
                  Price : ${{item.price}}
                </v-col>
              </v-row>
              <v-row dense class="mt-5 mb-5 mr-5 d-flex flex-row-reverse">
                  ${{ordertotal}}
              </v-row>
            </v-container> -->
          <v-container>
            <v-simple-table>
              <template v-slot:default>
                <thead>
                  <tr>
                    <!-- <th></th> -->
                    <th class="text-left">Name</th>
                    <th class="text-left">Price</th>
                    <th class="text-left">Qty</th>
                  </tr>
                </thead>
                <tbody>
                  <tr v-for="(item, index) in order" :key="index">
                    <!-- <td>
                      <v-hover color="red">
                        <v-btn color="red accent-1" elevation="2" dark small class="white--text px-0 py-0"
                          min-width="35" height="35" @click="$emit('delete', index)">
                          <v-icon x-small >mdi-delete</v-icon>
                        </v-btn>
                      </v-hover>
                    </td> -->
                    <td>{{ item.food_name | readmore}}</td>
                    <td>${{ item.price }}</td>
                    <td>{{ item.quantity}}</td>
                    <!-- <input type="text" :value="item.quantity"> -->
                    <!-- <td :value="item.quantity"></td> -->
                    <td style="height: 100%">
                      <div>
                        <v-icon slot="append" color="green" style="cursor: pointer"
                          @click="$emit('incrementOrder', index)">
                          mdi-plus
                        </v-icon>
                        <v-icon slot="prepend" color="red" style="cursor: pointer" @click="$emit('minusOrder', index)">
                          mdi-minus
                        </v-icon>
                      </div>
                    </td>
                    <td></td>

                  </tr>
                </tbody>
              </template>
            </v-simple-table>
            <v-row dense class="mt-5 mb-5 mr-5 d-flex flex-row-reverse">
              ${{ordertotal}}
            </v-row>
          </v-container>



          <v-divider></v-divider>

          <v-card-actions>
            <v-spacer></v-spacer>
            <!-- <v-btn color="primary" text @click="getdate">
              Test
            </v-btn> -->
            <v-btn color="primary" text @click="$emit('createOrder'),$emit('createBill')">
              <!-- <v-btn color="primary" text @click="createBill();createOrder()"> -->
              Order
            </v-btn>
            <v-btn color="red" text @click="$emit('empty')">
              Clear
            </v-btn>
            <v-btn color="primary" text @click="menuinfo = false">
              Close
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
  
    </v-app-bar>
  </div>
</template>

<script>
  import axios from 'axios';
  // import axios from 'axios';
  export default {
    props: ["order", "ordertotal", "dateid", "bill", "billtotal","t_no"],
    data: () => ({
      menuinfo: false,
      billinfo: false,
      reviewinfo: false,
      bill_id: "",
      review: "",
      current_name: '',
      current_bill_id: '',
      

    }),

    computed: {
      cartItemsCounter() {
        let counter = 0;
        this.order.forEach((value) => {
          if (value)
            counter++;
        })
        return counter;
      },
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
      }
    },
    mounted(){
      
    },


    methods: {
      getcurrent(foodname, bill_id) {
        this.current_name = foodname
        this.current_bill_id = bill_id
      },
      onCreatedReview() {
        const fd = ({
          food_name: this.current_name,
          bill_id: this.current_bill_id,
          reviews: this.review
        })
        const test = JSON.stringify(fd)
        console.log(test)
        axios.post(process.env.VUE_APP_BASE_URL+'/reviews/create-review', test)
          .then(res => {
            console.log(res)
          }).catch(error => {
            console.log(error)
          })
      }
      // createOrder() {
      //   const fd = JSON.stringify(this.order)
      //   // const bodyFormData = new FormData();
      //   // this.order.forEach((item) => {
      //   //   bodyFormData.append('order', item);
      //   // });
      //   // console.log(bodyFormData)
      //   // console.log("fd" + fd)
      //   axios.post('http://localhost:8081/bill/create-order', fd)
      //     .then(res => {
      //       console.log(res)
      //     }).catch(error => {
      //       console.log(error)
      //     })
      // },

      // createBill() {
      //   // const d = new Date();
      //   // const text = d.toISOString();
      //   this.bill_id = this.dateid
      //   const fdd = {
      //     "id": this.bill_id,
      //     "is_paid": false,
      //     "done": false,
      //     "table_no": 1,
      //     "amount": this.ordertotal,
      //   };
      //   // const test = JSON.stringify(fdd)

      //   // console.log("test :" + test)
      //   axios.post('http://localhost:8081/bill/create-bill', fdd)
      //     .then(res => {
      //       console.log(res)
      //     }).catch(error => {
      //       console.log(error)
      //     })
      // },

      // sendbillid(id) {
      //   this.$emit("sendid", id)
      // }
    }
  }
</script>

<style>

</style>