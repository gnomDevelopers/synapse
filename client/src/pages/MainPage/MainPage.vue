<template>
  <div class="min-h-screen bg-clr1 flex flex-col">
    
    <header class="h-16 bg-clr0 border-b border-border shadow-sm flex items-center justify-between px-6 sticky top-0 z-10">
      <div class="w-10 h-10 rounded-full overflow-hidden flex justify-center items-center cursor-pointer select-none" @click="$router.push({name:'MainPage'})">
        <img src="../../assets/images/logo.png"/>
      </div>


      <div class="flex items-center gap-3">
        <span class="text-gray-700 font-medium hidden sm:block">
          {{ user.fio }}
        </span>
        <UserAvatar :user-id="1" :size="'lg'"/>
      </div>
    </header>

    <RouterView />
  </div>
</template>
<script lang="ts">
import type { IUserResponse } from '@/api/api';
import UserAvatar from '@/widgets/PhotoTemplates/userAvatar.vue';
import { RouterView } from 'vue-router';

export default {
  components: {
    UserAvatar,
  },
  data() {
    return {
      user: {
        id: 1,
        fio: 'Иванов Иван Иванович',
      } as IUserResponse,
    }
  }
}
</script>