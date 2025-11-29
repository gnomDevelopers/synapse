<template>

  <div class="flex flex-col justify-start items-start relative">
    <label class="cursor-pointer" :for="'login-input-' + id">{{ label }}</label>
    <div class="login-input" :class="{'login-input-incorrect': errorMesssage !== ''}">
      <input 
      :type="(showPassword) ? ((isPasswordHide) ? type : 'text') : type" 
      class="grow cursor-pointer h-8" 
      :id="'login-input-' + id" 
      :placeholder="placeholder" 
      v-model="value" 
      @input="updateValue()"
      />

      <div v-if="showPassword" class="flex flex-col justify-center items-center cursor-pointer" @mousedown="isPasswordHide = false;" @mouseup="isPasswordHide = true;">
        <svg v-if="isPasswordHide" class="w-7 h-7 svg-cursor-pointer" viewBox="0 0 35 30" fill="none" xmlns="http://www.w3.org/2000/svg">
          <path d="M26.1625 22.425C23.6697 24.0538 20.6342 24.9561 17.5 25C7.29171 25 1.45837 15 1.45837 15C3.27238 12.1024 5.78836 9.57076 8.83754 7.575M14.4375 5.3C15.4414 5.0986 16.4691 4.99792 17.5 5C27.7084 5 33.5417 15 33.5417 15C32.6565 16.4195 31.6007 17.7559 30.3917 18.9875M20.5917 17.65C20.1912 18.0184 19.7082 18.3139 19.1715 18.5189C18.6348 18.7239 18.0555 18.8341 17.4681 18.843C16.8807 18.8518 16.2972 18.7592 15.7524 18.5706C15.2076 18.382 14.7128 18.1013 14.2973 17.7452C13.8819 17.3891 13.5544 16.9649 13.3343 16.498C13.1143 16.031 13.0062 15.5309 13.0166 15.0274C13.027 14.5239 13.1555 14.0273 13.3947 13.5673C13.6338 13.1073 13.9785 12.6933 14.4084 12.35M1.45837 1.25L33.5417 28.75" stroke="#0A035A" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"/>
        </svg>
        <svg v-else class="w-7 h-7 svg-cursor-pointer" viewBox="0 0 37 24" fill="none" xmlns="http://www.w3.org/2000/svg">
          <path d="M2 12C2 12 8 2 18.5 2C29 2 35 12 35 12" stroke="#0A035A" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"/>
          <path d="M2 12C2 12 8 22 18.5 22C29 22 35 12 35 12" stroke="#0A035A" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"/>
          <path d="M18.5 15.75C20.9853 15.75 23 14.0711 23 12C23 9.92893 20.9853 8.25 18.5 8.25C16.0147 8.25 14 9.92893 14 12C14 14.0711 16.0147 15.75 18.5 15.75Z" stroke="#0A035A" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"/>
        </svg>
      </div>
        
    </div>
    <div :class="errorMessageClass">
      <p class=" text-xs text-red-600">{{ errorMesssage }}</p>
    </div>
  </div>

</template>
<script lang="ts">
export default{
  emits: ['inputChange'],
  props:{
    id:{
      type: String,
      required: true,
    },
    type: {
      type: String,
      required: true,
    },
    placeholder: {
      type: String,
      required: true,
    },
    label: {
      type: String,
      required: true,
    },
    errorMesssage: {
      type: String,
      required: true,
    },
    showPassword: {
      type: Boolean,
      required: false,
      default: false,
    },
    defaultValue: {
      type: [Number, String],
      required: false,
      default: null
    },
    errorMessageClass: {
      type: String,
      required: false,
      default: 'h-4 select-none mt-1',
    },
    clearInput:{
      type: Boolean,
      required: false,
      default: false,
    }
  },
  data() {
    return{
      value: this.defaultValue,

      isPasswordHide: true,
    }
  },
  methods: {
    updateValue(){
      this.$emit('inputChange', this.value);
    }
  },
  watch:{
    //все ломает
    // defaultValue(newVal: number | string){
    //   this.value = newVal;
    // }
    clearInput(newVal: boolean){
      if(newVal){
        this.value = this.defaultValue;
      }
    }
  },
}
</script>
<style lang="css" scoped>
  input::-webkit-outer-spin-button,
  input::-webkit-inner-spin-button {
    -webkit-appearance: none;
    margin: 0;
  }
  .login-input{
    display: flex;
    flex-direction: row;
    flex-grow: 1;
    column-gap: 8px;
    width: 100%;
    padding: 4px 8px;
    font-size: 16px;
    border-radius: 6px;
    border: 2px solid var(--color-border);
    cursor: pointer;
    background-color: transparent;
  }
  .login-input:focus-within{
    outline: 2px solid var(--input-color-outline);
  }
  .login-input-incorrect{
    outline: 2px solid red;
  }
</style>