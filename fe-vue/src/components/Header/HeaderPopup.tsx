import { defineComponent } from "vue";
import { RouterLink } from "vue-router";

// Split comp for using at antother menu in future
const HeaderPopup = defineComponent({
    setup(props, { slots }) {
        return () => (
            <div class={"headerPopup"}>
                {slots.default?.()}
            </div>
        );
    },
});

type PopupColProps = {
    title: string,
    titleHref: string,
    listMenu: Array<{
        menu: string,
        menuHref: string,
    }>
}



export const HeaderPopupColContent = defineComponent({
    props: ["title", "titleHref", "listMenu"],
    setup(props: PopupColProps, { slots }) {
        const { title, titleHref, listMenu } = props
        return () => (
            <div class={"flex flex-col gap-2 "}>
                <RouterLink to={titleHref}>
                    <h5>{title}</h5>
                </RouterLink>
                {listMenu.map((x, i) => <RouterLink class={'header-link w-fit'} key={i} to={x.menuHref}>{x.menu}</RouterLink>)}
            </div>
        );
    },
});

export default HeaderPopup