function main() {
    const val = {
        abc: 'test'
    };

    add(val)

    console.log("val = ", val)
}

function add(val) {
    val.abc = "update"
}

main()