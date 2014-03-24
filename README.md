joker
=====

**joker** helps the caching using the `store_id` feature of `Squid
3.4` or the `storeurl_rewrite` feature of `Squid 2.7`.

It uses a *plugin* model, in which you can support new websites easily
without hassling too much with a long chain of `if else` code.

Setup
-----

I'm currently running on a OpenBSD 5.5-current machine, with the
squid-3.4.2p0 package.
On my `squid.conf` I have:

    store_id_program /usr/local/sbin/joker -new-format
    store_id_children 20 startup=0 idle=1 concurrency=100

    acl joker_access_list dstdomain i.imgur.com
    acl joker_access_list dstdomain .bp.blogspot.com .blogblog.com
    acl joker_access_list dstdomain .glbimg.com
    acl joker_access_list dstdomain .gravatar.com
    store_id_access allow joker_access_list
    store_id_access deny all

For now, the `-new-format` is required to enable the `store_id`
format, if not, it will be compatible with the `storeurl_rewrite`
feature. It supports concurrency, but I don't have a huge load of
traffic to test the optimal options.

ToDo
----

* Profile the code and identify any spots for improvement.
* Create a system to each plugin output his required squid acl.
* Create tests on the code.
* Select an appropriate license for the code.

Remarks
-------

I'm coding to distract myself from problems, or, for *fun*. I don't
have much time to support the code or to help, but feel free to mail
me, I'll try to give an reply as soon as possible.
