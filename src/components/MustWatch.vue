<template>
  <div class="section">
    <h2 class="section-title">必看视频</h2>
    
    <div class="videos-grid">
      <div 
        v-for="video in videos" 
        :key="video.id" 
        class="video-item"
      >
        <div class="video-thumbnail">
          <img 
            :src="video.thumbnail" 
            :alt="video.title"
            @error="handleImageError"
          />
          <div class="video-overlay">
            <div class="play-icon">▶</div>
          </div>
        </div>
        <div class="video-info">
          <h4 class="video-title">{{ video.title }}</h4>
          <p class="video-description">{{ video.description }}</p>
          <a 
            :href="video.url" 
            target="_blank" 
            rel="noopener noreferrer"
            class="video-link"
          >
            观看视频 →
          </a>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'MustWatch',
  data() {
    return {
      videos: [
        {
          id: 1,
          title: '解说精彩对局',
          description: '街霸6玩家高分对局 大司马本田VS毒药火舞',
          thumbnail: 'https://rolin-typora.oss-cn-guangzhou.aliyuncs.com/20250802235128.png',
          url: 'https://www.douyin.com/user/self?from_tab_name=main&modal_id=7514274830611385652'
        },
        {
          id: 2,
          title: '纯净精彩对局',
          description: '街霸6玩家顶分对局 阿光阿鬼VS不知火舞',
          thumbnail: 'https://rolin-typora.oss-cn-guangzhou.aliyuncs.com/20250802235310.png',
          url: 'https://www.douyin.com/user/self?from_tab_name=main&modal_id=7513618202321095946&showSubTab=compilation'
        },
        {
          id: 3,
          title: '超唐对局解说',
          description: '街霸6玩家超唐锦集,让你一次看个够',
          thumbnail: 'https://rolin-typora.oss-cn-guangzhou.aliyuncs.com/20250802235459.png',
          url: 'https://www.douyin.com/user/self?from_tab_name=main&modal_id=7524038804760579371&showSubTab=compilation'
        }
        // 后续可以通过硬编码添加更多视频
      ]
    }
  },
  methods: {
    handleImageError(event) {
      // 如果缩略图加载失败，使用默认图片
      event.target.src = 'data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMzAwIiBoZWlnaHQ9IjE2OCIgdmlld0JveD0iMCAwIDMwMCAxNjgiIGZpbGw9Im5vbmUiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyI+CjxyZWN0IHdpZHRoPSIzMDAiIGhlaWdodD0iMTY4IiBmaWxsPSIjZjhmOWZhIi8+CjxwYXRoIGQ9Ik0xMjAgODRDMTIwIDc1LjE2NDMgMTI3LjE2NCA2OCAxMzYgNjhIMTY0QzE3Mi44MzYgNjggMTgwIDc1LjE2NDMgMTgwIDg0VjEwMEMxODAgMTA4LjgzNiAxNzIuODM2IDExNiAxNjQgMTE2SDEzNkMxMjcuMTY0IDExNiAxMjAgMTA4LjgzNiAxMjAgMTAwVjg0WiIgZmlsbD0iI2NjY2NjYyIvPgo8cGF0aCBkPSJNMTUwIDg0TDE3MCA5NEwxNTAgMTA0Vjg0WiIgZmlsbD0iIzk5OTk5OSIvPgo8L3N2Zz4K'
    }
  }
}
</script>

<style scoped>
.videos-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 30px;
}

.video-item {
  border: 1px solid var(--border-color);
  border-radius: var(--border-radius);
  overflow: hidden;
  background-color: var(--background-color);
  transition: transform 0.3s ease, box-shadow 0.3s ease;
}

.video-item:hover {
  transform: translateY(-2px);
  box-shadow: var(--shadow);
}

.video-thumbnail {
  position: relative;
  width: 100%;
  height: 168px;
  overflow: hidden;
}

.video-thumbnail img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform 0.3s ease;
}

.video-item:hover .video-thumbnail img {
  transform: scale(1.05);
}

.video-overlay {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(0, 0, 0, 0.3);
  display: flex;
  justify-content: center;
  align-items: center;
  opacity: 0;
  transition: opacity 0.3s ease;
}

.video-item:hover .video-overlay {
  opacity: 1;
}

.play-icon {
  color: white;
  font-size: 48px;
  text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.5);
}

.video-info {
  padding: 20px;
}

.video-title {
  color: var(--primary-color);
  font-size: 18px;
  font-weight: 600;
  margin-bottom: 10px;
  line-height: 1.4;
}

.video-description {
  color: var(--text-color);
  font-size: 14px;
  line-height: 1.6;
  margin-bottom: 15px;
}

.video-link {
  display: inline-block;
  color: var(--accent-color);
  text-decoration: none;
  font-size: 14px;
  font-weight: 500;
  transition: color 0.3s ease;
}

.video-link:hover {
  color: #2980b9;
  text-decoration: underline;
}

@media (max-width: 768px) {
  .videos-grid {
    grid-template-columns: 1fr;
  }
  
  .video-thumbnail {
    height: 140px;
  }
}
</style> 