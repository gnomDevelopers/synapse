<template>
  <div class="flex flex-col gap-4 justify-start items-stretch">
    <loginInput 
      :id="'1'" 
      :type="'text'" 
      :placeholder="'Иванов Иван Иванович'" 
      :label="'ФИО:'" 
      :error-messsage="fioValid.error"
      :default-value="fioValid.value"
      @input-change="validFIOInput" />

    <loginInput 
      :id="'2'" 
      :type="'email'" 
      :placeholder="'example@synapse.ru'" 
      :label="'Электронная почта:'" 
      :error-messsage="emailValid.error"
      :default-value="emailValid.value"
      @input-change="validEmailInput" />
    
    <loginInput 
      :id="'3'" 
      :type="'password'" 
      :placeholder="'YourPassword123'" 
      :label="'Пароль:'" 
      :error-messsage="passwordValid.error"
      :show-password="true"
      :default-value="passwordValid.value"
      @input-change="validPasswordInput" />

    <loginInput 
      :id="'4'" 
      :type="'password'" 
      :placeholder="'YourPassword123'" 
      :label="'Повторите пароль:'" 
      :error-messsage="passwordRepeatValid.error"
      :show-password="true"
      :default-value="passwordRepeatValid.value"
      @input-change="validPasswordRepeatInput" />

    <input type="button" class="btn btn-main mt-8" value="Зарегистрироваться" @click="initLogIn"/> 
  </div>
</template>
<script lang="ts">

import { mapStores } from 'pinia';
import { useUserStore } from '@/stores/userStore';
import { ValidUserEmail, ValidUserPassword, ValidUserFIO } from '@/utils/validator';
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
      //валидаторы
      fioValid: {value: '', error: ''} as IValidator,
      emailValid: {value: '', error: ''} as IValidator,
      passwordValid: {value: '', error: ''} as IValidator,
      passwordRepeatValid: {value: '', error: ''} as IValidator,
    }
  },
  computed: {
    ...mapStores(useUserStore),
  },
  methods: {
    validEmailInput(value: string){
      this.emailValid = ValidUserEmail(value);
    },
    validPasswordInput(value: string){
      this.passwordValid = ValidUserPassword(value);

      if(this.passwordRepeatValid.value !== ''){ // проверка на совпадение паролей если в поле повторного пароля уже что-то введено
        this.validPasswordRepeatInput(this.passwordRepeatValid.value.toString());
      }
    },
    validPasswordRepeatInput(value: string){
      //если в осн пароле ошибка то сохраняем как есть
      if(this.passwordValid.error !== '') this.passwordRepeatValid = {value: value, error: ''};
      else {
        //если пароли совпадают то все ок
        if(this.passwordValid.value === value) this.passwordRepeatValid = {value: value, error: ''};
        //иначе сохраняем ввод но создаем ошибку
        else this.passwordRepeatValid = {value: value, error: 'Пароли не совпадают!'};
      }
    },
    validFIOInput(value: string){
      this.fioValid = ValidUserFIO(value);
    },
    initLogIn(){
      if(
        this.emailValid.value !== '' && 
        this.passwordValid.value !== '' &&
        this.passwordValid.value === this.passwordRepeatValid.value &&
        this.fioValid.value !== '') {

        const stID = this.StatusWindowAPI.createStatusWindow({status: this.StatusWindowAPI.getCodes.loading, text: 'Verifying', time: -1});
        
        // this.StatusWindowAPI.deleteStatusWindow(stID);
        // this.StatusWindowAPI.createStatusWindow({
        //   status: this.StatusWindowAPI.getCodes.success, 
        //   text: 'Профиль успешно создан! Ожидайте подтверждение учетной записи от администрации вашего ВУЗа', 
        //   time: 10000, 
        //   type: this.StatusWindowAPI.getTypes.detail
        // });

        // try{
        //   SET_COOKIE('access_token', login_res.data.access_token, new Date(Date.now() + 1000 * 60 * 60 * 24 * 30)); // +-2 минуты аксес токен
        //   SET_COOKIE('refresh_token', login_res.data.refresh_token, new Date(Date.now() + 1000 * 60 * 60 * 24 * 30)); // 30 дней рефреш токен
        // }
        // catch(e){ console.error(e);}

        // //Устанавливаем null чтобы при переходе роутер заново запросил информацию об аккаунте
        // this.userStore.isAuthorized = null;
      }

      if(this.fioValid.value === ''){
        return this.StatusWindowAPI.createStatusWindow({status: this.StatusWindowAPI.getCodes.error, text: 'Введите ФИО!'});
      }
      if(this.emailValid.value === ''){
        return this.StatusWindowAPI.createStatusWindow({status: this.StatusWindowAPI.getCodes.error, text: 'Введите email!'});
      }
      if(this.passwordValid.value === ''){
        return this.StatusWindowAPI.createStatusWindow({status: this.StatusWindowAPI.getCodes.error, text: 'Введите пароль!'});
      }
      if(this.passwordValid.value !== this.passwordRepeatValid.value){
        return this.StatusWindowAPI.createStatusWindow({status: this.StatusWindowAPI.getCodes.error, text: 'Пароли не совпадают!'});
      }
    }
  }
}
</script>