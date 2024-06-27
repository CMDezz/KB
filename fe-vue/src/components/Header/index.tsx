import { defineComponent, defineProps, normalizeProps, toRefs } from "vue";
// @ts-ignore
import { RouterLink } from "vue-router";
import './Header.css'
import LogoIcon from "@/assets/logo";
import CartIcon from "@/assets/cart";
import HeaderShopLink from "./HeaderShop";
import HeaderMobile from "./HeaderMobile";


type HeaderLinkProps = { title: string | Object, to: string }


export const HeaderLink = defineComponent({
    props: ['title', 'to'],
    render: (props: HeaderLinkProps) => {
        const { to, title } = props
        return (
            <RouterLink class={'header-link'} to={to}>
                <h5>{title}</h5>
            </RouterLink>

        )
    }
})

const Header = defineComponent({
    render: () => {
        return (
            <div id="Header" class={"sticky top-0 bg-white"}>
                <div class={"categoriesWrapper gap-2 md:gap-10"}>
                    <HeaderMobile class={"block md:hidden"} />
                    <HeaderLink to={'/'} title={<LogoIcon />} />
                    <HeaderShopLink class={"hidden md:block"} />
                    <HeaderLink class={"hidden md:block"} to={'/sale'} title={'SALE'} />

                </div>
                <div class={"actionsWrapper gap-5 md:gap-10"}>
                    <HeaderLink to={'/cart'} title={<CartIcon />} />
                    <HeaderLink to={'/cart'} title={'LOGIN'} />
                </div>
            </div>
        )
    }
})

export default Header