export interface ModuleProps {
  link: string
  text: string
  type: string
  icon: string
}

const modules: ModuleProps[] = [
  {
    text: '点位教学',
    link: '/',
    type: 'teach',
    icon: 'i-carbon:building-insights-1',
  },
  {
    text: '资讯',
    link: '/news',
    type: 'news',
    icon: 'i-ic:baseline-newspaper',
  },
  {
    text: '论坛',
    link: '/discuss',
    type: 'discuss',
    icon: 'i-clarity:talk-bubbles-line',
  },
]

export default modules
