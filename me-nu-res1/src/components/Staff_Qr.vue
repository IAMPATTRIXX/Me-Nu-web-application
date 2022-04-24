<template>
    <v-app>
        <Header />
        <v-main>
            <v-container>
                <div>
                    <StaffHeader />
                    <v-container fluid>
                        <v-row dense class="d-flex justify-end mt-10">
                            <v-btn @click="tableadd2">
                                Add table
                            </v-btn>
                        </v-row>
                        <v-row dense class="mt-5">
                            <v-col v-for="(data,i) in tables" :key="i" :cols="3">
                                <v-badge overlap  :color="data.status?'green':'red'">
                                    <v-card height="auto" width="auto">
                                    
                                        <v-row dense>
                                            <v-col cols=8>
                                                <v-card-title>Table No.{{data.table_no}}</v-card-title>
                                            </v-col>
                                            <v-col cols=3>
                                                <v-switch v-model="data.status" class="pl-5"
                                                    @change="changeStatus(data.status,data.table_no,data.id)">
                                                </v-switch>
                                            </v-col>
                                        </v-row>
                                        <v-card-text>
                                            <v-row dense>
                                                <v-col cols="6">
                                                
                                                    <v-btn @click="createdurl(data.table_no)" elevation="2">
                                                        <v-icon class="pr-2">fas fa-qrcode</v-icon>
                                                        Gen QR
                                                    </v-btn>
                                                </v-col>
                                                <v-col cols="6">
                                                    <v-btn @click="gopath(data.table_no)" class="ml-5">
                                                        <v-icon class="pr-2">fas fa-square-up-right</v-icon>
                                                        Go
                                                    </v-btn>
                                                </v-col>
                                                
                                            </v-row>
                                        </v-card-text>
                                    </v-card>

                                </v-badge>
                            </v-col>
                        </v-row>
                        <span> {{this.url}} </span>
                    </v-container>
                </div>
            </v-container>
        </v-main>
    </v-app>
</template>

<script>
    import StaffHeader from '@/components/Staff_Header.vue'
    import Header from '@/components/HeaderView'
    import axios from 'axios'

    export default {
        components: {
            StaffHeader,
            Header
        },
        data: () => ({
            url: '',
            tables: [],
            tabless: [],
            count: 0,
            counttest: 1,
            status: false,
            urll: [],
            t_id:""

        }),
        mounted() {
        },
        created() {
            // setInterval(() => {
            this.getdata()
            // },1000)
            // console.log("created_out"+this.tables.length)
        },


        methods: {
            changeStatus(status, tNo) {
                console.log(status)

                this.status = status
                const fd = {
                    "table_no": tNo,
                    // "id": id,
                    // "url": this.urll[tNo - 1],
                    "status": this.status
                }
                const jsondata = JSON.stringify(fd)
                console.log(jsondata)
                axios.post(process.env.VUE_APP_BASE_URL+'/table/update-table', jsondata)
                    .then(res => {
                        console.log('success')
                        console.log(res)
                    }).catch(error => {
                        console.log(error)
                    })

            },

            createdurl(tID) {
                
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
                let time = days.concat(months, years, hours, mins, secs);
                let url = "https://me-nu-cus.web.app/" //ต้องเปลี่ยน URL ของ Menu ที่ดีพลอยแล้ว
                let test = url.concat(`${tID}`, "/", time)
                this.urll[tID - 1] = test
                // console.log("url ของโต๊ะที่ " + tID + this.urll[tID - 1])

               
                const fd = {
                    "table_no": tID,
                    // "id": tID,
                    "url": this.urll[tID - 1],
                    "status": this.status,
                    "t_key": time
                }
                const jsondata = JSON.stringify(fd)
                axios.post(process.env.VUE_APP_BASE_URL+'/table/update-table', jsondata)
                    .then(res => {
                        console.log(res)
                    }).catch(error => {
                        console.log(error)
                    })
               
            },

            tableadd2() {
                
                if (this.tables.length == 0) {
                    localStorage.setItem("count", 1)
                    var a = parseInt(localStorage.getItem("count"))
                    const fd = {
                        "table_no": a,
                        "status": false,
                    };
                    const test = JSON.stringify(fd)
                    console.log(test)
                    axios.post(process.env.VUE_APP_BASE_URL+'/table/create-table', test)
                        .then(res => {
                            console.log(res)
                            this.getdata()
                        }).catch(error => {
                            console.log(error)
                        })
                    localStorage.setItem("count", a + 1)
                } else if (this.tables.length != 0) {
                    var b = parseInt(localStorage.getItem("count"))
                    const fd = {
                        "table_no": b,
                        "status": false,
                    };
                    const test = JSON.stringify(fd)
                    console.log(test)
                    axios.post(process.env.VUE_APP_BASE_URL+'/table/create-table', test)
                        .then(res => {
                            console.log(res)
                            this.getdata()
                        }).catch(error => {
                            console.log(error)
                        })
                    localStorage.setItem("count", b + 1)
                }
            },

            gopath(tID) {
                this.$router.push({
                    path: `/staff/showqr/?tID=${tID}`
                })
                // router.push({ path: `/user/${userId}` })
            },
            getdata(){
                axios
                .get(process.env.VUE_APP_BASE_URL+'/table/get-table')
                .then(res => {
                    this.tables = res.data.data.item.table
                })
            }
        },
    }
</script>


<style>

</style>