import MenuIcon from "@/assets/menu";
import MenuClose from "@/assets/menuClose";
import { clickOutside } from "@/directives/clickOutside";
import { defineComponent, ref } from "vue";


const HeaderMobilePopup = defineComponent({
    props: ["toggleOpen"],
    directives: { clickOutside: clickOutside },
    render(props: { toggleOpen: () => void }) {
        const isFirstTime = ref(true)
        return <div class={" overlay bg-neutral"}>
            <div v-click-outside={() => {
                if (isFirstTime.value) {
                    isFirstTime.value = false
                    return;
                }
                props.toggleOpen()

            }} class={" HeaderMobilePopup"}>
                <p>heheh</p>
            </div>
        </div>
    }
})


const HeaderMobile = defineComponent({
    setup() {
        const isOpen = ref(false)

        const toggleOpen = () => {
            isOpen.value = !isOpen.value
        }
        return {
            isOpen,
            toggleOpen
        }
    },
    render() {
        const { isOpen, toggleOpen } = this
        return <div class={" HeaderMobile"}>
            <button onClick={toggleOpen}>
                {
                    isOpen ?
                        <MenuClose /> :
                        <MenuIcon />
                }
            </button>
            {isOpen && <HeaderMobilePopup toggleOpen={toggleOpen} />}
        </div>
    }
})

export default HeaderMobile