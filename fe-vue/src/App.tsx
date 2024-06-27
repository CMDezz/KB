import { defineComponent } from 'vue'
import { RouterView } from 'vue-router' // Import RouterView from vue-router

import "@/assets/css/base.css"
import "@/assets/css/reset.css"
import "@/assets/css/font.css"

import router from './router'
import Header from './components/Header'
import Footer from './components/Footer'
export default defineComponent({
    setup: () => {
        return {
            router
        }
    },
    render: () => {
        return (
            <div id="app">
                <Header />
                <div id='app-content'>
                    <RouterView />
                </div>
                <Footer />
            </div>
        )
    }
})