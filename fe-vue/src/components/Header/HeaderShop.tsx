import { defineComponent } from "vue";
import { HeaderLink } from ".";
import HeaderPopup, { HeaderPopupColContent } from "./HeaderPopup";

type HeaderLinkProps = { title: string | Object, to: string }



const HeaderShopLink = defineComponent({
    props: [],
    render: (props: HeaderLinkProps) => {
        return (
            <div class={"headerShopWrapper headerPopupContainer"}>
                <HeaderLink to={'/shop'} title={'SHOP'} />
                <HeaderPopup class={"flex gap-x-10 2xl:gap-20 lg:gap-15"}>
                    <HeaderPopupColContent
                        title={'MECHANIC KEYBOARD'}
                        titleHref={'/'}
                        listMenu={[{ menu: "Shop all", menuHref: "/" },
                        { menu: "Keyboards", menuHref: "/" },
                        { menu: "Switchs", menuHref: "/" },
                        { menu: "Accessories", menuHref: "/" }
                        ]}
                    />
                    <HeaderPopupColContent
                        title={'KEYCAPS'}
                        titleHref={'/'}
                        listMenu={[{ menu: "Shop all", menuHref: "/" },
                        { menu: "Keycaps", menuHref: "/" },
                        { menu: "Artisan Keycaps", menuHref: "/" },
                        { menu: "Cloned Keycaps", menuHref: "/" }
                        ]}
                    />
                    <HeaderPopupColContent
                        title={'BATTLE STATIONS'}
                        titleHref={'/'}
                        listMenu={[{ menu: "Shop all", menuHref: "/" },
                        { menu: "Desk Mats", menuHref: "/" },
                        { menu: "Lighting", menuHref: "/" },
                        ]}
                    />

                    <HeaderPopupColContent
                        title={'COLLECTIONS'}
                        titleHref={'/'}
                        listMenu={[{ menu: "Shop all", menuHref: "/" },
                        { menu: "Drop Collapse", menuHref: "/" },
                        ]}
                    />
                </HeaderPopup>
            </div>
        )
    }
})



export default HeaderShopLink