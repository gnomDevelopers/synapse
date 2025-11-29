<template>
  <div class="w-svw h-svh flex flex-row">
    <div class="login-bg-left h-full w-0 lg:w-1/2 overflow-hidden relative">
      <!--image background-->
      <div class="absolute w-full h-full bg-main/68"></div>
      <div class="absolute w-full h-full backdrop-blur-xs"></div>
      <div class="absolute inset-0 z-10 flex flex-col justify-center items-center"> <!--bg-main-->
        <img src="../../assets/images/auth-page-logo.png" class="w-[400px]"/>
      </div>
    </div>
    <div class="h-full w-full lg:w-1/2 flex flex-col justify-center items-center login-bg-right relative">
      <div class="absolute w-full h-full bg-main/68 lg:hidden"></div>
      <div class="absolute w-full h-full backdrop-blur-xs lg:hidden"></div>
      <div class="absolute p-4 sm:p-10 rounded-xl w-auto lg:w-full xl:w-3/4 2xl:w-3/5 flex flex-col justify-start items-stretch gap-8 bg-clr0">
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

import loginSwitcher from './features/authSwitch.vue';

export default {
  components:{
    loginSwitcher,
  },
  methods: {
    redirectByAuthType(index: number) {
      switch (index) {
        case 0: this.$router.push({name: "AuthPageLogin"}); break;
        case 1: this.$router.push({name: "AuthPageSignUp"}); break;
        default: this.$router.push({name: "AuthPageLogin"}); break;
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
      background-color: var(--color-clr0);
      background-image: none;
    }
  }
</style>