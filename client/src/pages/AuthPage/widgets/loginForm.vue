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
import { login } from '@/api/api';

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
    async initLogIn(){
      if(this.emailValid.value !== '' && this.passwordValid.value !== '') {
        const stID = this.StatusWindowAPI.createStatusWindow({status: this.StatusWindowAPI.getCodes.loading, text: 'Verifying', time: -1});
        

        try {
          const res = await login({
            email: this.emailValid.value,
            password: this.passwordValid.value,
          });

          this.StatusWindowAPI.deleteStatusWindow(stID);

          try{
            SET_COOKIE('access_token', res.data.access_token, new Date(Date.now() + 1000 * 60 * 60 * 24 * 30));
            SET_COOKIE('refresh_token', res.data.refresh_token, new Date(Date.now() + 1000 * 60 * 60 * 24 * 30));
          }
          catch(e){ console.error(e);}

          this.userStore.isAuthorized = true;
          this.userStore.id = 15;
          this.userStore.fio = 'Unknown';

          this.$router.push({name: "MainPage"});
        } catch(e) {
          this.StatusWindowAPI.deleteStatusWindow(stID);
          this.StatusWindowAPI.createStatusWindow({status: this.StatusWindowAPI.getCodes.error, text: 'Verify failed', time: 2000});
        }
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