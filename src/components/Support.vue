<template>
  <div class="section">
    <h2 class="section-title">支持主包</h2>
    
    <div class="support-content">
      <div class="support-intro">
        <p>您的支持是我持续创作优质内容的动力！</p>
        <ul>
          <li>🎬 帮助我创作更好的视频内容</li>
          <li>💪 解决我的经济压力，让我更专注于创作</li>
          <li>🚀 支持我购买更好的设备和软件</li>
          <li>❤️ 让我能够持续为社区贡献价值</li>
        </ul>
      </div>

      <div class="support-methods">
        <div class="support-item">
          <h4>微信赞赏</h4>
          <div class="support-image" @click="showImageModal('wechat')">
            <img 
              :src="supportImages.wechat"
              alt="微信赞赏码"
              @error="handleImageError"
            />
          </div>
          <p>微信扫码赞赏</p>
        </div>

        <div class="support-item">
          <h4>支付宝</h4>
          <div class="support-image" @click="showImageModal('alipay')">
            <img 
              :src="supportImages.alipay"
              alt="支付宝收款码"
              @error="handleImageError"
            />
          </div>
          <p>支付宝扫码支持</p>
        </div>
      </div>
    </div>

    <!-- 图片模态框 -->
    <ImageModal
      :visible="modalVisible"
      :image-src="modalImageSrc"
      :image-alt="modalImageAlt"
      @close="closeImageModal"
    />
  </div>
</template>

<script>
import ImageModal from './ImageModal.vue'
// 正确导入图片资源，确保Vite能够处理
import wechatSupportImage from '/src/assets/images/wechat-support.jpg'
import alipaySupportImage from '/src/assets/images/alipay-support.jpg'

export default {
  name: 'Support',
  components: {
    ImageModal
  },
  data() {
    return {
      modalVisible: false,
      modalImageSrc: '',
      modalImageAlt: '',
      supportImages: {
        wechat: wechatSupportImage,
        alipay: alipaySupportImage
      }
    }
  },
  methods: {
    showImageModal(type) {
      this.modalImageSrc = this.supportImages[type]
      this.modalImageAlt = type === 'wechat' ? '微信赞赏码' : '支付宝收款码'
      this.modalVisible = true
    },

    closeImageModal() {
      this.modalVisible = false
    },

    handleImageError(event) {
      // 如果图片加载失败，使用默认图片
      event.target.src = 'data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMjAwIiBoZWlnaHQ9IjIwMCIgdmlld0JveD0iMCAwIDIwMCAyMDAiIGZpbGw9Im5vbmUiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyI+CjxyZWN0IHdpZHRoPSIyMDAiIGhlaWdodD0iMjAwIiBmaWxsPSIjZjhmOWZhIi8+CjxwYXRoIGQ9Ik04MCAxMDBDODAgODguOTU0MyA4OC45NTQzIDgwIDEwMCA4MEMxMTEuMDQ2IDgwIDEyMCA4OC45NTQzIDEyMCAxMDBDMTIwIDExMS4wNDYgMTExLjA0NiAxMjAgMTAwIDEyMEM4OC45NTQzIDEyMCA4MCAxMTEuMDQ2IDgwIDEwMFoiIGZpbGw9IiNjY2NjY2MiLz4KPHN2ZyB3aWR0aD0iNDAiIGhlaWdodD0iNDAiIHZpZXdCb3g9IjAgMCA0MCA0MCIgZmlsbD0ibm9uZSIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4KPHBhdGggZD0iTTIwIDIwTDI4IDI4TDIwIDM2TDEyIDI4TDIwIDIwWiIgZmlsbD0iIzk5OTk5OSIvPgo8L3N2Zz4KPC9zdmc+'
    }
  }
}
</script>

<style scoped>
.support-content {
  display: flex;
  flex-direction: column;
  gap: 30px;
}

.support-intro {
  background-color: var(--light-background);
  padding: 20px;
  border-radius: var(--border-radius);
  border-left: 4px solid var(--accent-color);
}

.support-intro p {
  margin-bottom: 15px;
  font-size: 16px;
  color: var(--text-color);
}

.support-intro ul {
  list-style: none;
  padding-left: 0;
}

.support-intro li {
  margin-bottom: 8px;
  padding-left: 20px;
  position: relative;
  font-size: 14px;
  color: var(--text-color);
}

.support-intro li::before {
  content: '';
  position: absolute;
  left: 0;
  top: 8px;
  width: 6px;
  height: 6px;
  background-color: var(--accent-color);
  border-radius: 50%;
}

.support-methods {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 30px;
}

.support-item {
  text-align: center;
  padding: 20px;
  border: 1px solid var(--border-color);
  border-radius: var(--border-radius);
  background-color: var(--background-color);
  transition: transform 0.3s ease, box-shadow 0.3s ease;
}

.support-item:hover {
  transform: translateY(-2px);
  box-shadow: var(--shadow);
}

.support-item h4 {
  color: var(--primary-color);
  margin-bottom: 15px;
  font-size: 18px;
}

.support-image {
  cursor: pointer;
  margin-bottom: 10px;
  border-radius: var(--border-radius);
  overflow: hidden;
  box-shadow: var(--shadow);
  transition: transform 0.3s ease;
}

.support-image:hover {
  transform: scale(1.05);
}

.support-image img {
  width: 100%;
  height: auto;
  display: block;
}

.support-item p {
  color: var(--light-text);
  font-size: 14px;
  margin: 0;
}

@media (max-width: 768px) {
  .support-methods {
    grid-template-columns: 1fr;
  }
}
</style> 