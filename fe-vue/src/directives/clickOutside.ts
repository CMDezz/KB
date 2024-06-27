import type { DirectiveBinding } from 'vue';

interface ClickOutsideElement extends HTMLElement {
    clickOutsideEvent?: (event: Event) => void;
}

export const clickOutside = {
    beforeMount(element: ClickOutsideElement, binding: DirectiveBinding) {
        console.log({
            element,
            binding
        });

        // Check that click was outside the element and its children
        element.clickOutsideEvent = function (event: Event) {
            // Call method provided in the attribute value if click was outside
            if (!(element === event.target || element.contains(event.target as Node))) {
                binding.value?.(event);
            }
        };
        document.body.addEventListener('click', element.clickOutsideEvent);
    },
    unmounted(element: ClickOutsideElement) {
        document.body.removeEventListener('click', element.clickOutsideEvent as any);
    }
};
