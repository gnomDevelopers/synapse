<template>
  <div 
    class="rounded-full overflow-hidden shrink-0"
    :class="sizeClasses"
  >
    <img 
      :src="avatarSrc" 
      alt="User Avatar" 
      class="w-full h-full object-cover"
    />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

const ASSETS_PATH = '../../assets/images/avatars';

export default defineComponent({
  name: 'UserAvatar',
  props: {
    userId: {
      type: Number,
      required: true,
      default: 0
    },
    size: {
      type: String, // например: 'sm', 'md', 'lg'
      default: 'md'
    }
  },
  computed: {
    avatarIndex(): number {
      return (Math.abs(this.userId) % 13) + 1;
    },
    avatarSrc(): string {
      try {
        return new URL(`${ASSETS_PATH}/avatar-${this.avatarIndex}.png`, import.meta.url).href;
      } catch (e) {
        console.error('Ошибка загрузки аватара:', e);
        return new URL(`${ASSETS_PATH}/avatar-1.png`, import.meta.url).href;
      }
    },
    sizeClasses(): string {
      switch (this.size) {
        case 'sm': return 'w-8 h-8';   // 32px
        case 'md': return 'w-10 h-10'; // 40px
        case 'lg': return 'w-12 h-12'; // 48px
        case 'xl': return 'w-16 h-16'; // 64px
        case '2xl': return 'w-24 h-24'; // 96px
        default: return 'w-10 h-10';
      }
    }
  }
});
</script>