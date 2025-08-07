<template>
  <div class="section">
    <h2 class="section-title">交流群</h2>
    
    <div class="community-content">
      <div class="community-intro">
        <p>加入我们的交流群，与志同道合的格斗游戏爱好者一起交流！</p>
        <ul>
          <li>🎮 方便约对局，找到合适的对手</li>
          <li>💬 实时交流游戏技巧和心得</li>
          <li>👥 认识更多格斗游戏朋友</li>
          <li>📢 第一时间获取最新视频更新</li>
        </ul>
      </div>

      <div class="qr-codes">
        <div class="qr-item">
          <h4>微信群</h4>
          <div class="qr-image" @click="showImageModal('wechat')">
            <img 
              :src="qrImages.wechat"
              alt="微信群二维码"
              @error="handleImageError"
            />
          </div>
          <p>扫码加入微信群</p>
        </div>

        <div class="qr-item">
          <h4>QQ群</h4>
          <div class="qr-image" @click="showImageModal('qq')">
            <img 
              :src="qrImages.qq"
              alt="QQ群二维码"
              @error="handleImageError"
            />
          </div>
          <p>扫码加入QQ群</p>
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
import wechatQrImage from '/src/assets/images/wechat-qr.jpg'
import qqQrImage from '/src/assets/images/qq-qr.jpg'

export default {
  name: 'Community',
  components: {
    ImageModal
  },
  data() {
    return {
      modalVisible: false,
      modalImageSrc: '',
      modalImageAlt: '',
      qrImages: {
        wechat: wechatQrImage,
        qq: qqQrImage
      }
    }
  },
  methods: {
    showImageModal(type) {
      this.modalImageSrc = this.qrImages[type]
      this.modalImageAlt = type === 'wechat' ? '微信群二维码' : 'QQ群二维码'
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
.community-content {
  display: flex;
  flex-direction: column;
  gap: 30px;
}

.community-intro {
  background-color: var(--light-background);
  padding: 20px;
  border-radius: var(--border-radius);
  border-left: 4px solid var(--accent-color);
}

.community-intro p {
  margin-bottom: 15px;
  font-size: 16px;
  color: var(--text-color);
}

.community-intro ul {
  list-style: none;
  padding-left: 0;
}

.community-intro li {
  margin-bottom: 8px;
  padding-left: 20px;
  position: relative;
  font-size: 14px;
  color: var(--text-color);
}

.community-intro li::before {
  content: '';
  position: absolute;
  left: 0;
  top: 8px;
  width: 6px;
  height: 6px;
  background-color: var(--accent-color);
  border-radius: 50%;
}

.qr-codes {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 30px;
}

.qr-item {
  text-align: center;
  padding: 20px;
  border: 1px solid var(--border-color);
  border-radius: var(--border-radius);
  background-color: var(--background-color);
  transition: transform 0.3s ease, box-shadow 0.3s ease;
}

.qr-item:hover {
  transform: translateY(-2px);
  box-shadow: var(--shadow);
}

.qr-item h4 {
  color: var(--primary-color);
  margin-bottom: 15px;
  font-size: 18px;
}

.qr-image {
  cursor: pointer;
  margin-bottom: 10px;
  border-radius: var(--border-radius);
  overflow: hidden;
  box-shadow: var(--shadow);
  transition: transform 0.3s ease;
}

.qr-image:hover {
  transform: scale(1.05);
}

.qr-image img {
  width: 100%;
  height: auto;
  display: block;
}

.qr-item p {
  color: var(--light-text);
  font-size: 14px;
  margin: 0;
}

@media (max-width: 768px) {
  .qr-codes {
    grid-template-columns: 1fr;
  }
}
</style> 