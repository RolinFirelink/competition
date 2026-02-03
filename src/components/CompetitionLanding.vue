<template>
  <div class="landing-page">
    <header class="hero">
      <div class="brand">Street Fighter 6 教学中心</div>
      <div class="hero-actions">
        <router-link class="pill active" to="/competition">赛事中心</router-link>
        <button class="pill disabled" disabled>攻略中心 · 敬请期待</button>
        <button class="pill disabled" disabled>直播与录播 · 敬请期待</button>
      </div>
      <p class="tagline">从入门到进阶，赛事、教学一站式体验</p>
    </header>

    <section class="ads-strip" aria-label="推广位">
      <div class="ads-track" :style="trackStyle">
        <div
          v-for="(ad, index) in adsLoop"
          :key="index"
          class="ad-item"
        >
          <a :href="ad.link" target="_blank" rel="noopener noreferrer">
            <img :src="ad.image" :alt="ad.title" />
            <span>{{ ad.title }}</span>
          </a>
        </div>
      </div>
    </section>

    <main class="content">
      <div class="section-title">快速入口</div>
      <div class="options-grid">
        <router-link class="option-card" to="/competition">
          <div class="option-label">赛事中心</div>
          <div class="option-desc">杯赛记录 · 报名 · 赛程</div>
        </router-link>
        <div class="option-card disabled">
          <div class="option-label">教学攻略</div>
          <div class="option-desc">角色教学 · 连招指令 · 对局心得</div>
        </div>
        <div class="option-card disabled">
          <div class="option-label">直播/录像</div>
          <div class="option-desc">赛事直播 · 高光集锦</div>
        </div>
      </div>
    </main>
  </div>
</template>

<script>
export default {
  name: 'CompetitionLanding',
  data() {
    return {
      ads: [
        {
          title: '盖世小鸡 启明星手柄',
          link: 'https://www.gamestar.com',
          image:
            'https://dummyimage.com/320x120/1f1f1f/ffffff&text=Sponsor+1'
        },
        {
          title: '每日一灯 赛事直播',
          link: 'https://www.bilibili.com',
          image:
            'https://dummyimage.com/320x120/222222/ffffff&text=Sponsor+2'
        },
        {
          title: '街霸6 玩家社区',
          link: 'https://www.capcom.com',
          image:
            'https://dummyimage.com/320x120/262626/ffffff&text=Sponsor+3'
        }
      ],
      speed: 40
    };
  },
  computed: {
    adsLoop() {
      return [...this.ads, ...this.ads];
    },
    trackStyle() {
      const duration = (this.ads.length * 10) / (this.speed / 40);
      return {
        '--duration': `${duration}s`
      };
    }
  }
};
</script>

<style scoped>
.landing-page {
  min-height: 100vh;
  padding: 56px 32px 72px;
  color: #e7ecf5;
  font-family: 'Inter', 'Segoe UI', system-ui, -apple-system, sans-serif;
  background:
    radial-gradient(circle at 20% 20%, rgba(255, 99, 164, 0.14), transparent 22%),
    radial-gradient(circle at 80% 10%, rgba(115, 174, 255, 0.14), transparent 20%),
    radial-gradient(circle at 50% 85%, rgba(116, 255, 199, 0.08), transparent 24%),
    linear-gradient(135deg, #0b0c14 0%, #070810 100%);
}

.hero {
  max-width: 1200px;
  margin: 0 auto 32px auto;
  text-align: center;
}

.brand {
  font-size: 34px;
  font-weight: 800;
  letter-spacing: 1.2px;
  margin-bottom: 16px;
  color: #ffffff;
  text-transform: uppercase;
  text-shadow: 0 8px 32px rgba(0, 0, 0, 0.45);
}

.hero-actions {
  display: flex;
  gap: 12px;
  justify-content: center;
  flex-wrap: wrap;
  margin-bottom: 18px;
}

.pill {
  border: 1px solid rgba(255, 255, 255, 0.14);
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.08), rgba(255, 255, 255, 0.02));
  color: #cfd7e6;
  padding: 12px 22px;
  border-radius: 999px;
  font-weight: 700;
  letter-spacing: 0.4px;
  text-transform: uppercase;
  cursor: pointer;
  transition: all 0.2s ease;
  box-shadow: 0 10px 40px rgba(0, 0, 0, 0.25);
}

.pill.active {
  border-color: rgba(255, 99, 164, 0.55);
  color: #ffffff;
  box-shadow:
    0 0 0 1px rgba(255, 99, 164, 0.25),
    0 14px 48px rgba(255, 99, 164, 0.35);
  background: linear-gradient(135deg, #ff5c93, #6c7bff);
}

.pill.disabled {
  cursor: not-allowed;
  opacity: 0.55;
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.02), rgba(255, 255, 255, 0.01));
}

.tagline {
  margin-top: 6px;
  color: #b8c3d9;
  font-weight: 600;
  font-size: 15px;
}

.ads-strip {
  margin: 36px auto 44px;
  max-width: 1120px;
  overflow: hidden;
  border-radius: 18px;
  border: 1px solid rgba(255, 255, 255, 0.08);
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.05), rgba(255, 255, 255, 0.015));
  box-shadow:
    0 20px 80px rgba(0, 0, 0, 0.4),
    inset 0 1px 0 rgba(255, 255, 255, 0.06);
  padding: 6px 0;
}

.ads-track {
  display: flex;
  gap: 14px;
  animation: scroll var(--duration) linear infinite;
  width: max-content;
  padding: 4px 8px;
}

.ad-item {
  min-width: 320px;
  background: linear-gradient(135deg, rgba(26, 28, 36, 0.95), rgba(16, 17, 24, 0.9));
  border-radius: 12px;
  border: 1px solid rgba(255, 255, 255, 0.06);
  overflow: hidden;
  box-shadow: 0 6px 20px rgba(0, 0, 0, 0.32);
  transition: transform 0.2s ease, border-color 0.2s ease;
}

.ad-item:hover {
  transform: translateY(-4px);
  border-color: rgba(255, 99, 164, 0.4);
}

.ad-item a {
  display: flex;
  flex-direction: column;
  color: inherit;
  text-decoration: none;
}

.ad-item img {
  width: 100%;
  height: 126px;
  object-fit: cover;
  background: #0f1116;
  filter: saturate(1.08);
}

.ad-item span {
  padding: 11px 14px;
  font-weight: 700;
  font-size: 14px;
  color: #dde4f6;
}

@keyframes scroll {
  from {
    transform: translateX(0);
  }
  to {
    transform: translateX(calc(-50% - 7px));
  }
}

.content {
  max-width: 1120px;
  margin: 0 auto;
}

.section-title {
  font-size: 18px;
  font-weight: 800;
  color: #f4f6ff;
  letter-spacing: 0.5px;
  margin-bottom: 16px;
  text-transform: uppercase;
}

.options-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(240px, 1fr));
  gap: 20px;
}

.option-card {
  background: linear-gradient(145deg, rgba(255, 92, 147, 0.16), rgba(108, 123, 255, 0.14)),
    linear-gradient(180deg, rgba(255, 255, 255, 0.05), rgba(255, 255, 255, 0.015)),
    rgba(12, 13, 20, 0.94);
  border: 1px solid rgba(255, 255, 255, 0.12);
  border-radius: 16px;
  padding: 18px;
  color: #e9edf5;
  text-decoration: none;
  display: flex;
  flex-direction: column;
  gap: 8px;
  transition: transform 0.18s ease, border-color 0.2s ease, box-shadow 0.2s ease;
  box-shadow: 0 18px 50px rgba(0, 0, 0, 0.42);
  position: relative;
  overflow: hidden;
}

.option-card::after {
  content: '';
  position: absolute;
  inset: 0;
  background: radial-gradient(circle at 20% 20%, rgba(255, 255, 255, 0.07), transparent 30%);
  pointer-events: none;
}

.option-card:hover {
  transform: translateY(-6px);
  border-color: rgba(255, 255, 255, 0.2);
  box-shadow: 0 18px 60px rgba(255, 92, 147, 0.18);
}

.option-card.disabled {
  cursor: not-allowed;
  opacity: 0.58;
  border-style: dashed;
  box-shadow: none;
}

.option-label {
  font-size: 18px;
  font-weight: 800;
  letter-spacing: 0.4px;
  color: #ffffff;
}

.option-desc {
  font-size: 13px;
  color: #c5cee0;
}

@media (max-width: 720px) {
  .landing-page {
    padding: 36px 20px 56px;
    background: linear-gradient(135deg, #0b0c14 0%, #070810 100%);
  }

  .brand {
    font-size: 26px;
  }

  .ad-item {
    min-width: 260px;
  }
}
</style>

