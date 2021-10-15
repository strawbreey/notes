---
title: "NgZoneRun"
date: 2021-10-09T11:01:45+08:00
draft: false
---

### Angular2 zone.run() vs ChangeDetectorRef.detectChanges()

ApplicationRef.tick (same as setTimeout()), and zone.run() cause change detection on the whole application. Also event listeners added within Angular or by Angular (using view bindings or @HostBinding() cause change detection for the whole application.

ChangeDetectorRef.detectChanges runs change detection for a specific component (and its descendants if applicable, for example because of input bindings)

If some code running outside Angular's zone calls into Angular's code and changes the state, then change detection needs to be invoked explicitly because Angular has no way of knowing that the state changed.

If the change to the state is local to a component (for example a components field), ChangeDetectorRef.detectChanges or ChangeDetectorRef.markforCheck are more efficient.

If the call from outside for example navigates to a different route, this can have consequences to a number of components, and it's also not clear when the whole route change is completed because it might cause async calls (and callbacks being called). In this case zone.run() is the better option, because the code directly and indirectly invoked (like callbacks of observables and promises) invoked will run inside Angular's zone and Angular will recognize them and invoke change detection automatically.

### 参考资料


- https://stackoverflow.com/questions/51455545/when-to-use-ngzone-run