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

    <!-- 语音频道部分 -->
    <div class="voice-section">
      <h3 class="voice-title">🎤 语音频道</h3>
      <div class="voice-content">
        <div class="voice-info">
          <p>应部分粉丝要求现已新增了语音频道，有需要的可以加，方便约战时进入语音</p>
        </div>
        
        <div class="voice-platform">
          <div class="platform-header">
            <h4>🎮 OOPZ语音</h4>
            <span class="platform-badge">推荐</span>
          </div>
          <div class="platform-details">
            <div class="detail-item">
              <span class="detail-label">语音链接：</span>
              <a href="https://oopz.cn/i/G3WuB2" target="_blank" class="voice-link">https://oopz.cn/i/G3WuB2</a>
            </div>
            <div class="detail-item">
              <span class="detail-label">频道ID：</span>
              <span class="channel-id">016079069</span>
            </div>
            <div class="detail-item">
              <span class="detail-label">使用说明：</span>
              <span class="usage-text">用浏览器找oopz或者在手机电脑上下载oopz软件，寻找域用ID</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 俱乐部部分 -->
    <div class="club-section">
      <h3 class="club-title">🏆 游戏俱乐部</h3>
      <div class="club-content">
        <div class="club-info">
          <p>目前我们有俱乐部，名字是LightClub，直接搜索并加入就好了，我并没有设置检验，加入俱乐部可以方便约战</p>
        </div>
        
        <div class="club-details">
          <div class="club-card">
            <div class="club-icon">🎯</div>
            <div class="club-info-text">
              <h4>LightClub</h4>
              <p>格斗游戏爱好者聚集地</p>
              <div class="club-features">
                <span class="feature-tag">约战</span>
                <span class="feature-tag">交流</span>
                <span class="feature-tag">技巧分享</span>
              </div>
            </div>
            <div class="join-club">
              <button class="btn btn-join" @click="copyClubName">
                📋 复制俱乐部名称
              </button>
            </div>
          </div>
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
    },

    copyClubName() {
      navigator.clipboard.writeText('LightClub').then(() => {
        // 可以添加一个提示消息
        alert('俱乐部名称 "LightClub" 已复制到剪贴板！')
      }).catch(err => {
        console.error('复制失败:', err)
        // 降级方案
        const textArea = document.createElement('textarea')
        textArea.value = 'LightClub'
        document.body.appendChild(textArea)
        textArea.select()
        document.execCommand('copy')
        document.body.removeChild(textArea)
        alert('俱乐部名称 "LightClub" 已复制到剪贴板！')
      })
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

/* 语音频道样式 */
.voice-section {
  margin-top: 40px;
  padding: 25px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-radius: var(--border-radius);
  color: white;
}

.voice-title {
  font-size: 22px;
  font-weight: 600;
  margin-bottom: 20px;
  display: flex;
  align-items: center;
  gap: 10px;
}

.voice-content {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.voice-info p {
  font-size: 16px;
  line-height: 1.6;
  opacity: 0.95;
}

.voice-platform {
  background: rgba(255, 255, 255, 0.1);
  border-radius: var(--border-radius);
  padding: 20px;
  backdrop-filter: blur(10px);
}

.platform-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 15px;
}

.platform-header h4 {
  font-size: 18px;
  font-weight: 600;
  margin: 0;
}

.platform-badge {
  background: #ff6b6b;
  color: white;
  padding: 4px 12px;
  border-radius: 20px;
  font-size: 12px;
  font-weight: 500;
}

.platform-details {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.detail-item {
  display: flex;
  align-items: center;
  gap: 10px;
}

.detail-label {
  font-weight: 500;
  min-width: 80px;
  opacity: 0.9;
}

.voice-link {
  color: #ffd93d;
  text-decoration: none;
  word-break: break-all;
  transition: color 0.3s ease;
}

.voice-link:hover {
  color: #fff;
  text-decoration: underline;
}

.channel-id {
  background: rgba(255, 255, 255, 0.2);
  padding: 4px 12px;
  border-radius: 6px;
  font-family: 'Courier New', monospace;
  font-weight: 600;
}

.usage-text {
  opacity: 0.9;
  line-height: 1.5;
}

/* 俱乐部样式 */
.club-section {
  margin-top: 40px;
  padding: 25px;
  background: linear-gradient(135deg, #f093fb 0%, #f5576c 100%);
  border-radius: var(--border-radius);
  color: white;
}

.club-title {
  font-size: 22px;
  font-weight: 600;
  margin-bottom: 20px;
  display: flex;
  align-items: center;
  gap: 10px;
}

.club-content {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.club-info p {
  font-size: 16px;
  line-height: 1.6;
  opacity: 0.95;
}

.club-details {
  background: rgba(255, 255, 255, 0.1);
  border-radius: var(--border-radius);
  padding: 20px;
  backdrop-filter: blur(10px);
}

.club-card {
  display: flex;
  align-items: center;
  gap: 20px;
  flex-wrap: wrap;
}

.club-icon {
  font-size: 48px;
  background: rgba(255, 255, 255, 0.2);
  border-radius: 50%;
  width: 80px;
  height: 80px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.club-info-text {
  flex: 1;
  min-width: 200px;
}

.club-info-text h4 {
  font-size: 20px;
  font-weight: 600;
  margin-bottom: 8px;
}

.club-info-text p {
  margin-bottom: 12px;
  opacity: 0.9;
}

.club-features {
  display: flex;
  gap: 8px;
  flex-wrap: wrap;
}

.feature-tag {
  background: rgba(255, 255, 255, 0.2);
  padding: 4px 12px;
  border-radius: 20px;
  font-size: 12px;
  font-weight: 500;
}

.join-club {
  flex-shrink: 0;
}

.btn-join {
  background: rgba(255, 255, 255, 0.2);
  border: 2px solid rgba(255, 255, 255, 0.3);
  color: white;
  transition: all 0.3s ease;
}

.btn-join:hover {
  background: rgba(255, 255, 255, 0.3);
  border-color: rgba(255, 255, 255, 0.5);
  transform: translateY(-2px);
}

@media (max-width: 768px) {
  .qr-codes {
    grid-template-columns: 1fr;
  }
  
  .club-card {
    flex-direction: column;
    text-align: center;
  }
  
  .club-icon {
    width: 60px;
    height: 60px;
    font-size: 36px;
  }
  
  .voice-section,
  .club-section {
    padding: 20px;
  }
  
  .detail-item {
    flex-direction: column;
    align-items: flex-start;
    gap: 5px;
  }
  
  .detail-label {
    min-width: auto;
  }
}
</style> 