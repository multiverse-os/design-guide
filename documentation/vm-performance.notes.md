# Controller VM preformance tweaking
Geting some jittering with the new setup, its capable of running high end games but there may be some final tweaking required.

- JS is ridiculous, it blocks the UI of its own websites, its getting used more than ever, its super dangerous because of XSS, it seemingly has no limit on resources usage and people who program it are often novices who are not concerned about resource usage, or at least finish their project and need to be complained to before it even crosses their mind thats a problem to consider. They may never discover it relates to power usage or more importantly: battery drain. This is obvious considering the trend has been to push logic off the server and into the JS, meaning mobile devices use their battery faster than they need to if the design was more concious. That said: javascript usage can be confined and controlled in their amount of total resources allowed creating a hard limit on their ability to drag down a whole client computer. 

* Huge pages could still be used to squeeze out more power

* In the same way the CPUs are isolated, if we can isolate the memory used by the Host machine from the memory of the Controller and two router VMs. Then you remove another major vector for breakout attacks.

* A UI to select CPU features, name the CPU, and control it would be wundervoll, maybe even a full UI showing the virtual computer. Letting the user remove features like power button so for example the host cant turn it off or on. Clicking on the parts then being able to switch out the features. Letting people understand they are basically able to program the physical aspects of their computer.

* Control over cpufreq, maybe like when certain apps are turned on known to have profiles of larger usage, change the settings. When they are turned off, fallback. THis can save energy.

* Reflash bios on every boot, keep stored bios in a hidden drive that has read only access.

* The point is to extend the concept of multiple hop proxies, like with onion routing in Tor. Except with computers in a network, forcing a proxy between, each USB drive, between each Ethernet card, between any direct access to the Controller VM. Allowing ephemeral machines to function as proxies, so if they become infected, the infection is deleted when they are reset.

* The controller should have custom BIOS, custom CPU. Exposing only secure features needed, leaving the rest. Then a custom kernel, only providing features needed.

* Take steps (including testing with hypervisor detection tools) to make the VM undecteable and the HOST not detectable as a hypervisor. Includes (1) change MAC on network cards, (2) Change information on virtual components

* I found a scientific article explaining why Xen clocks are better than other VM software. Then they explained how to make an even better one than Xen. Implement this clock in Multiverse.

* ufw like tool for port forwding, opening ports on VMs, doing tunneling, multihop connecting

* The controller VM should not have any applications, anything that has not been established as a permentant Application VM, should open up in a disposable VM. The disposable VM would be configured similarly to the controller VM, so one could modify files and such, but the system files would all be thrown away when the X was hit. So any program like terminal or anything else.

* UI to visualize the networks. 

* auto wifi hacking- use this: https://github.com/xtr4nge/FruityWifi as a reference, kinda lame its php, make it better 

* [cool unique idea] Arm, and other micro controller VMs, built to mimic standard devices. Possible to insert sensors and so on, capable of feeding it random data. Simplyfing the process of programming robots and make the process cheaper. Once the code is done, you can buy all the pieces you need and have the code ready to run on the assembled bot.
  Also androids and other mobiesl

### Controller VM

Controller VM is setup to get the PCI graphics card the usb slots and so on. It also gets isolated CPUs. 

This not only vastly improves the preformance but it also lets you receive massive security benefits not seen in other VM/Container based compartmentalization operating systems. By telling the Host machine kernel not to schedule the Cores dedicated to the controller and by ensuring you correctly paired the hyperthreads when doing pinning. You ensure that the host and the controller use separate CPUs removing one of the biggest vectors of VM breakout. Exploits that take advantage of running code on the same CPU as the host and leveraging their shared access will not work shrinking the attack surface massively. 

One may ask why even use VMs, the primary reason is because we create a logical shim between the hardware and the virtual machine. Allowing programmatic control over BIOS flashing/updating/configuration, over CPU (you can just the CPUID features you want, leave out dangerous ones (watching CVE you can remove these making your CPU secure despite running a chip with known problems. like leave out management features, but keep all preformance features.) This promatic control over a baremetal machine would require robotic control to do everything one can with Multieverwse controller VM. This fine grain control provides a completely different user experience, now one does not need to wait for Intel, and gets more power than Intel in deciding what features their CPU has and is not forced to use the features sold. This is a powerful concept and fundamentally changes the security of general use computing.

Attacks against the machine would most likely infect the VM BIOS/firmware and not the host. And the BIOS can literally be reflashed on every boot, providing unpresented security. By walling off the HOST machine, ripping out everything extra and giving it no control over the Controller VM, malware in the host BIOS should have difficulty in attacking Multiverse OS.



## Utility VM Ideas
* more general but wget for whole folder?

* Backup server

* Interal DNS, ensure an onion request never leaves the network so it can be accidently sent over the internet and reveal dns information in the failed lookup

* USB proxy vms

* Electronic disobedience, easy set and forget DOS tools. Should apply attacks and check effect, use variety of tools, spit out all the recon it can do automatically and suggest attacks. Routes traffic through multihops to hide origin.

* Tor Non-Exit/Exit Relay

* Automatically convert pdfs to open formats
  - https://github.com/Debian/gpdftext


* Home web hosting, linking up with friends and even freinds of friends from contacts DB. If any site gets a ton of usage, people can share their resources and provide additional hosts, scaling up when use is larger and sclaing down when its not needed. This way people host form home but they also get the benefits of sclaing up in cloud centers. Lets actually ahve the internet not just be hosted by 1 provider creating a giant central point of weakness and essentially breaking the entire concept that makes the internet good

* [important to have early] setup to provide packages for OS upgrade of the host machine

https://github.com/strothj/alpine-debtool 


* Reproducible build VM SERVER not fucking container. Continious integration testing and so on.

## Key Ideas

* No logs on any computer in the cluster BUT the controller, all logs should be output to this server so its all centralized.
https://github.com/clustree/journalbeat-deb
 - https://github.com/mbenkmann/garcon
 - maybe a middleware for caddy!

# https://github.com/systemd/mkosi
**This is critical, in Multiverse, we want to build all the isos for the App and Utility VMs. This would allow for custom building of images. If this is not used it should be the guide. It even supports adding luks and other stuff to the iamge and multiple images.**

Unlike most bullshit like docker or even Qubes, we want the Apps to be distributed WITHOUT fucking prebuilt images that you have to trust. Removal of all trust is critical to security. So the images should be built from auditable buildscripts

*This may be even just useful for building images for VMs inside of Multiverse because it can make a wide variety of custom iso images.*

A fancy wrapper around dnf --installroot, debootstrap, pacstrap and zypper that may generate disk images with a number of bells and whistles.

In theory, any distribution may be used on the host for building images containing any other distribution, as long as the necessary tools are available. Specifically, any distro that packages debootstrap may be used to build Debian or Ubuntu images. Any distro that packages dnf may be used to build Fedora images. Any distro that packages pacstrap may be used to build Arch Linux images. Any distro that packages zypper may be used to build openSUSE images.

Additionally, bootable GPT disk images (as created with the --bootable flag) work when booted directly by EFI systems, for example in KVM via:

*Generated images are legacy-free. This means only GPT disk labels (and no MBR disk labels) are supported, and only systemd based images may be generated. Moreover, for bootable images only EFI systems are supported (not plain MBR/BIOS).*

qemu-kvm -m 512 -smp 2 -bios /usr/share/edk2/ovmf/OVMF_CODE.fd -drive format=raw,file=image.raw

# Multiverse OS Controller VM Performance
Multiverse OS is capable of having Controller VM performance near bare-metal speeds. After doing the below performance tuning, Multiverse OS developers were able to play modern FPS (Deus Ex: Human Revolution and Stellaris were both tested with perfect results) on their dedicated gaming Controller VMs.

There are even potential paths for development that will lead to further performance increases on all Multiverse OS virtual machines, especially Controller VMs, including but not limited to: memory-based screen relaying of Application VMs to the Controller VM, custom Multiverse OS tcp/ip network stack to avoid interaction with host machine kernelspace, and other planned improvements to the Multiverse OS system. 

The configuration changes discussed below will be applied by editing the libvirt XML:

```
virsh edit debian9.controller.multiverse
```

### Performance Tuning
Performance tuning gains come primarily from two main topics: (1) CPU Pinning, where we specify explicitly what CPU cores are dedicated to the VM, along with explicitly defining what CPU features we want passed through to the Controller VM (which enhances both security and performance); and (2) Clock/Tick configuration, where we define how the VM CPU should handle clock/ticks in a way that does not affect other VM performance. 

In addition to the above topics there are other enhancements that will be listed but the main boosts come from the above two configuration changes. This includes exposing `l3-cache` if the bare-metal host machine hardware supports it and other similar changes.

#### CPU Pinning
CPU Pinning is where we define explicitly what CPUs are configured to be used by the VM along with explicitly define what CPU features will be passed.

A key concept with CPU pinning is ensuring that correctly paired CPU cores are passed together. Because of `hyperthreading` provided by essentially all modern CPUs, each core has an associated pair, and very importantly the cores are not simply paired (1,2), (3,4), (5,6)... 

The physical core mapping can be found in the data output from `cat /proc/cpuinfo` but to save time the following command can be used to cut out just the relevant information:

```
paste <(cat /proc/cpuinfo | grep "core id") <(cat /proc/cpuinfo | grep "processor") | sort
```
The result will show the `core id` number along with the `processor` number, below is the output for an `i5` intel processor: 

```
core id : 0 processor : 0 
core id : 1 processor : 1
core id : 2 processor : 2
core id : 3 processor : 3
core id : 0 processor : 4
core id : 1 processor : 5
core id : 2 processor : 6
core id : 3 processor : 7
```

On this example `i5` machine the pairs look like:

```
# The first core is made up of two sub-cores numbered
0,4
# The second core is made up of two sub-cores numbered
1,5
# The third core is made of up two sub-cores numbered
2,6
# The fourth core is made of up two sub-cores numbered
3,7
```
When CPU pinning for Multiverse with the above `i5` we use three physical cores and their associated processors and leave the last for the routers and the host machine to use. Grouping by physical cores is the most efficient, otherwise it would pair cores in different physical locations causing lag. 

This results in the following XML configuration, placed below the `<currentMemory>` element. Note in this instance we are giving the VM two composite vCPUs (vCPU 6 and 7) that can use any CPU in the cpuset specified. This is optional, and may help with performance of VMs on machines that don't have many physical cores.

```
<vcpu placement='static' cpuset='1-3,5-7'>8</vcpu>
<cputune>
  <vcpupin vcpu='0' cpuset='3'/>
  <vcpupin vcpu='1' cpuset='7'/>
  <vcpupin vcpu='2' cpuset='1'/>
  <vcpupin vcpu='3' cpuset='5'/>
  <vcpupin vcpu='4' cpuset='2'/>
  <vcpupin vcpu='5' cpuset='6'/>
  <vcpupin vcpu='6' cpuset='1-3,5-7'/>
  <vcpupin vcpu='7' cpuset='1-3,5-7'/>
</cputune>
```

CPU Pinning is where Multiverse OS developers found the most significant performance boost outside of GPU passthrough and is an essential part of the Multiverse OS experience. 

#### Clock/Tick Configuration
The next boost comes from correctly configuring clocks/ticks. Eventually Multiverse OS developers plan to implement a customized version of the method used by `Xen` VMs to synchronize/update ticks, but until this can be completed, the following is the best configuration for Qemu/KVM clocks. The key point here is QEMU avoiding delays and correcting clock skew after the fact in combination with a script run on the guest VM. 

Below is optimized for the same example `Intel i5` used in the previous section:

```
<clock offset='utc'>
  <timer name='rtc' tickpolicy='catchup' track='guest'>
    <catchup threshold='123' slew='120' limit='10000'/>
  </timer>
  <timer name='pit' tickpolicy='discard'/>
  <timer name='hpet' present='no'/>
  <timer name='hypervclock' present='no'/>
</clock>
```

Multiverse OS developers always encourage learning more about the topic and refining the configuration to your device specifications. It is always better to understand configuration options over mindlessly setting them.

#### CPU Passthrough
One key step in setting up a Multiverse OS Controller is configuring CPU `host-passthrough` which allows us to explicitly define what CPU features we want passed to the controller VM. This allows for the Multiverse OS CPU to essentially be a custom set of CPU features, enabling us to not pass through CPU features that are potential security risks and only passing through features that are known to be secure and add performance. The net result is that this step increases both the performance and security of our Multiverse OS Controller VM. 

Below is optimized for the `Intel i5`:

```
<cpu mode='host-passthrough'>
  <topology sockets='1' cores='4' threads='2'/>
  <feature policy='disable' name='lahf_lm'/>
  <feature policy='require' name='fpu'/>
  <feature policy='require' name='pse'/>
  <feature policy='require' name='pse36'/>
  <feature policy='require' name='bmi2'/>
  <feature policy='require' name='rtm'/>
  <feature policy='require' name='lm'/>
  <feature policy='require' name='avx2'/>
  <feature policy='require' name='apic'/>
  <feature policy='require' name='mmx'/>
  <feature policy='require' name='aes'/>
  <feature policy='require' name='nx'/>
  <feature policy='require' name='pdpe1gb'/>
  <feature policy='require' name='clflush'/>
  <feature policy='require' name='vme'/>
  <feature policy='require' name='ss'/>
  <feature policy='require' name='avx'/>
  <feature policy='require' name='hle'/>
  <feature policy='require' name='erms'/>
  <feature policy='require' name='xsave'/>
  <feature policy='require' name='hypervisor'/>
  <feature policy='require' name='cx16'/>
  <feature policy='require' name='popcnt'/>
  <feature policy='require' name='movbe'/>
  <feature policy='require' name='sse2'/>
  <feature policy='require' name='ssse3'/>
  <feature policy='require' name='sse4.1'/>
  <feature policy='require' name='sse4.2'/>
</cpu>
```

#### Adding Features
Below are features added on an `i5` cpu to increase performance, review available features for your processor and add accordingly.

```
<features>
  ...
  <hap state='on'/>
  <pmu state='off'/>
  ...
</features>
```

**NOTE** The Multiverse OS developer who determined these features useful notes are mixed in with the old notes, so these need to be explained thoroughly so we are not mindlessly adding features we do not know. Everyone should always be fully informed what they are doing when configuring Multiverse OS machines.

#### Enabling L3 Cache
If your CPU supports L3 cache, you will benefit significantly by enabling it in the Multiverse OS controller VM. This is done by adding custom Qemu command-line arguments which requires the use of a custom schema.

The schema is changed by modifying the first `<domain>` element to:

```
<domain type='kvm' xmlns:qemu='http://libvirt.org/schemas/domain/qemu/1.0'>
```

Unfortunately they do not provide secure `https` access, and are required to only access the schema via `http`.

This enables the use of the `<qemu:command-line>` tag used to pass custom command-line arguments to Qemu when launching the VM.

**DEVELOPMENT** If Multiverse OS provides its own schema it can be served over `https` to improve the security and prevent tampering with schemas when downloading. 

The above can not be saved without adding the following, otherwise it will automatically be removed.

```
<qemu:commandline>
  <qemu:arg value='-cpu'/>
  <qemu:arg value='host,l3-cache=on'/>
</qemu:commandline>
```

