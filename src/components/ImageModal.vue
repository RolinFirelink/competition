<template>
  <div v-if="visible" class="modal" @click="closeModal">
    <div class="close-btn" @click="closeModal">&times;</div>
    <div class="modal-content" @click.stop>
      <img :src="imageSrc" :alt="imageAlt" />
    </div>
  </div>
</template>

<script>
export default {
  name: 'ImageModal',
  props: {
    visible: {
      type: Boolean,
      default: false
    },
    imageSrc: {
      type: String,
      required: true
    },
    imageAlt: {
      type: String,
      default: '图片'
    }
  },
  methods: {
    closeModal() {
      this.$emit('close')
    }
  }
}
</script>

<style scoped>
.modal {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.8);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
  backdrop-filter: blur(5px);
}

.close-btn {
  position: absolute;
  top: 20px;
  right: 30px;
  font-size: 40px;
  color: white;
  cursor: pointer;
  background: rgba(0, 0, 0, 0.5);
  width: 50px;
  height: 50px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: background-color 0.3s ease;
  z-index: 1001;
}

.close-btn:hover {
  background: rgba(0, 0, 0, 0.8);
}

.modal-content {
  max-width: 90%;
  max-height: 90%;
  position: relative;
  border-radius: 8px;
  overflow: hidden;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.5);
}

.modal-content img {
  width: 100%;
  height: 100%;
  object-fit: contain;
  display: block;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .modal-content {
    max-width: 95%;
    max-height: 95%;
  }
  
  .close-btn {
    top: 10px;
    right: 15px;
    font-size: 30px;
    width: 40px;
    height: 40px;
  }
}
</style> 