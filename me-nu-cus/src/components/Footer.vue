<template>
<v-container class="mt-10">
    <v-footer dark padless absolute>
        <v-card class="flex" flat tile>
            <!-- <v-card-title class="teal">
                <strong class="subheading">Me-nU!</strong>

                <v-spacer></v-spacer>

                <v-btn v-for="icon in icons" :key="icon" class="mx-4" dark icon>
                    <v-icon size="24px">
                        {{ icon }}
                    </v-icon>
                </v-btn>
            </v-card-title> -->

            <v-card-text class="py-2 white--text text-center">
                {{ new Date().getFullYear() }} — <strong>Me-nU</strong>
            </v-card-text>
        </v-card>
    </v-footer>

</v-container>
</template>

<script>
    export default {
        data: () => ({
            icons: [
                'mdi-facebook',
                'mdi-twitter',
                'mdi-linkedin',
                'mdi-instagram',
            ],
        }),
    }
</script>