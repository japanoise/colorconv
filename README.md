# colorconv

Convert konsole colorschemes into st colorschemes.

It's not a terribly clever conversion, so make sure to give it a once-over
before you use it.

Once you're sure it's good, run:

    export SCHEME="some-color-scheme"
    colorconv ~/.local/share/konsole/$SCHEME.colorscheme | xsel --clipboard

and paste into config.h in the relevant section.
