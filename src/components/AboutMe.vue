<template>
  <div class="section">
    <h2 class="section-title">关于我</h2>
    <div class="about-content">
      <div class="avatar-container">
        <img 
          :src="avatarImage" 
          alt="每日一灯头像" 
          class="avatar"
          @error="handleImageError"
        />
      </div>
      <div class="profile-info">
        <h3 class="profile-name">每日一灯</h3>
        <p class="profile-signature">格斗游戏爱好者，专注分享精彩对局</p>
        <div class="profile-description">
          <p>我是一名专注于格斗游戏视频解说、IT开发和街霸6视频剪辑分享的内容创作者。</p>
          <p>主要领域包括：</p>
          <ul>
            <li>🎮 格斗游戏视频解说</li>
            <li>💻 IT开发技术分享</li>
            <li>🎬 街霸6视频剪辑</li>
            <li>🏆 大师对局分析</li>
          </ul>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
// 正确导入图片资源，确保Vite能够处理
import avatarImage from '/src/assets/images/avatar.jpg'

export default {
  name: 'AboutMe',
  data() {
    return {
      avatarImage: avatarImage
    }
  },
  methods: {
    handleImageError(event) {
      // 如果头像加载失败，使用默认头像
      event.target.src = 'data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMTIwIiBoZWlnaHQ9IjEyMCIgdmlld0JveD0iMCAwIDEyMCAxMjAiIGZpbGw9Im5vbmUiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyI+CjxyZWN0IHdpZHRoPSIxMjAiIGhlaWdodD0iMTIwIiByeD0iNjAiIGZpbGw9IiMzNDk4ZGIiLz4KPGNpcmNsZSBjeD0iNjAiIGN5PSI0NSIgcj0iMjAiIGZpbGw9IndoaXRlIi8+CjxwYXRoIGQ9Ik0yMCAxMDBDNDAgNzAgODAgNzAgMTAwIDEwMEgyMFoiIGZpbGw9IndoaXRlIi8+Cjwvc3ZnPgo='
    }
  }
}
</script>

<style scoped>
.about-content {
  display: flex;
  align-items: flex-start;
  gap: 30px;
}

.avatar-container {
  flex-shrink: 0;
}

.profile-info {
  flex: 1;
}

.profile-name {
  font-size: 28px;
  font-weight: 600;
  color: var(--primary-color);
  margin-bottom: 10px;
}

.profile-signature {
  font-size: 16px;
  color: var(--accent-color);
  font-style: italic;
  margin-bottom: 20px;
}

.profile-description {
  line-height: 1.8;
}

.profile-description p {
  margin-bottom: 15px;
}

.profile-description ul {
  list-style: none;
  padding-left: 0;
}

.profile-description li {
  margin-bottom: 8px;
  padding-left: 20px;
  position: relative;
}

.profile-description li::before {
  content: '';
  position: absolute;
  left: 0;
  top: 8px;
  width: 6px;
  height: 6px;
  background-color: var(--accent-color);
  border-radius: 50%;
}

@media (max-width: 768px) {
  .about-content {
    flex-direction: column;
    text-align: center;
  }
  
  .avatar-container {
    align-self: center;
  }
}
</style> 