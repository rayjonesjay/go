window.addEventListener("load", function () {
    const copyButton = document.querySelector('#copy-text');
    const resultPre = document.querySelector('#result-pre');
    const tooltip = document.querySelector('.tooltip');

    let hide = true;
    resultPre.addEventListener('mouseenter', () => {
        copyButton.style.display = 'block';
    })

    resultPre.addEventListener('mouseleave', () => {
        setTimeout(() => {
            if (hide) {
                copyButton.style.display = 'none';
            }
        }, 5);
    })

    copyButton.addEventListener('mouseenter', () => {
        hide = false;
        copyButton.style.display = 'block';
    });

    copyButton.addEventListener('mouseleave', () => {
        hide = true;
    });

    copyButton.addEventListener('click', () => {
        navigator.clipboard.writeText(resultPre.innerText).then(r => {
            tooltip.style.display = 'block';
            setTimeout(() => {
                tooltip.style.display = 'none';
            }, 3000);
        });
    });
});
