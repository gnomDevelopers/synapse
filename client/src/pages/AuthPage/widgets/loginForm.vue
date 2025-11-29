<template>
  <div class="flex flex-col gap-4 justify-start items-stretch">
    <loginInput 
      :id="'1'" 
      :type="'text'" 
      :placeholder="'example@synapse.ru'" 
      :label="'Электронная почта:'" 
      :error-messsage="emailValid.error"
      @input-change="validEmailInput" />

    <loginInput 
      :id="'2'" 
      :type="'password'" 
      :placeholder="'YourPassword123'" 
      :label="'Пароль:'" 
      :error-messsage="passwordValid.error"
      :show-password="true"
      @input-change="validPasswordInput" />
  </div>
  <input type="button" class="btn btn-main mt-8" value="Войти" @click="initLogIn"/>
</template>
<script lang="ts">

import { mapStores } from 'pinia';
import { useUserStore } from '@/stores/userStore';
import { ValidUserEmail, ValidUserPassword } from '@/utils/validator';
import { type IValidator } from '@/utils/constants';
import { useStatusWindowAPI } from '@/widgets/StatusWindow/statusWindowAPI';

import loginInput from '@/pages/AuthPage/features/authInput.vue';
import { SET_COOKIE } from '@/utils/functions';

export default {
  components:{
    loginInput,
  },
  data(){
    return {
      StatusWindowAPI: useStatusWindowAPI(),

      emailValid: {value: '', error: ''} as IValidator,
      passwordValid: {value: '', error: ''} as IValidator,
    }
  },
  computed:{
    ...mapStores(useUserStore),
  },
  methods: {
    validEmailInput(value: string){
      this.emailValid = ValidUserEmail(value);
    },
    validPasswordInput(value: string){
      this.passwordValid = ValidUserPassword(value);
    },
    initLogIn(){
      if(this.emailValid.value !== '' && this.passwordValid.value !== '') {
        const stID = this.StatusWindowAPI.createStatusWindow({status: this.StatusWindowAPI.getCodes.loading, text: 'Verifying', time: -1});
        
        // const body: ILogIn = {
        //   email:    this.emailValid.value     as string,
        //   password: this.passwordValid.value  as string,
        // };

        // API_LogIn(body)
        // .then((res: any) => {
        //   this.StatusWindowAPI.deleteStatusWindow(stID);
        //   this.StatusWindowAPI.createStatusWindow({status: this.StatusWindowAPI.getCodes.success, text: 'Вы успешно вошли в аккаунт!'});

        //   try{
        //     SET_COOKIE('access_token', res.data.access_token, new Date(Date.now() + 1000 * 60 * 60 * 24 * 30)); // +-2 минуты аксес токен
        //     SET_COOKIE('refresh_token', res.data.refresh_token, new Date(Date.now() + 1000 * 60 * 60 * 24 * 30)); // 30 дней рефреш токен
        //   }
        //   catch(e){ console.error(e);}

        //   //Устанавливаем null чтобы при переходе роутер заново запросил информацию об аккаунте
        //   this.userStore.isAuthorized = null;

        //   this.$router.push({name: 'MainPage'});
        // })
        // .catch((err) => {
        //   this.StatusWindowAPI.deleteStatusWindow(stID);
        //   this.StatusWindowAPI.createStatusWindow({status: this.StatusWindowAPI.getCodes.error, text: 'Пользователь не найден!'});
        // })
      }

      if(this.emailValid.value === ''){
        return this.StatusWindowAPI.createStatusWindow({status: this.StatusWindowAPI.getCodes.error, text: 'Введите email!'});
      }
      if(this.passwordValid.value === ''){
        return this.StatusWindowAPI.createStatusWindow({status: this.StatusWindowAPI.getCodes.error, text: 'Введите пароль!'});
      }
    }
  }
}
</script>