<template>
  <div class="w-svw h-svh flex flex-row">
    <div class="login-bg-left h-full w-0 lg:w-1/2 ">
      <!--image background-->
    </div>
    <div class="h-full w-full lg:w-1/2 flex flex-col justify-center items-center login-bg-right bg-clr1">
      <div class=" p-4 sm:p-10 rounded-xl w-auto lg:w-full xl:w-3/4 2xl:w-3/5 flex flex-col justify-start items-stretch gap-8 ">
        <h2 class="text-center text-3xl font-light">Добро пожаловать в Synapse!</h2>

          <loginSwitcher 
            :headers-array="['Вход', 'Регистрация']" 
            :selected-default="0" 
            class="mb-8"
            @header-select="redirectByAuthType" 
          />

          <RouterView />

      </div>
    </div>
  </div>
</template>
<script lang="ts">

import loginSwitcher from './features/loginSwitcher.vue';

export default {
  components:{
    loginSwitcher,
  },
  methods: {
    redirectByAuthType(index: number) {
      switch (index) {
        case 0: this.$router.push({name: ""}); break;
        case 1: this.$router.push({name: ""}); break;
        default: this.$router.push({name: ""}); break;
      }
    }
  }
}
</script>
<style lang="css" scoped>
  .login-bg-left{
    background-image: none;
  }
  .login-bg-right{
    --blur-color: rgba(39, 151, 39, 0.5);
    background-image: url('@/assets/images/login-bg.png');
    background-repeat: no-repeat;
    background-size: cover;
    background-position: center;
  }
  @media(min-width: 1024px){
    .login-bg-left{
      background-image: url('@/assets/images/login-bg.png');
      background-repeat: no-repeat;
      background-size: cover;
      background-position: center;
    }
    .login-bg-right{
      background-color: var(--bg-color-0);
      background-image: none;
    }
  }
</style>