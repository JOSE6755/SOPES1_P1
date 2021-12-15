#include <linux/module.h>
#include <linux/kernel.h>
#include <linux/init.h>
#include <linux/proc_fs.h>
#include <asm/uaccess.h>
#include <linux/seq_file.h>
#include <linux/hugetlb.h>

MODULE_LICENSE("GPL");
MODULE_DESCRIPTION("Modulo de RAM SOPES 1");
MODULE_AUTHOR("Jose Martinez");

struct sysinfo info;

static int escribir_archivo(struct seq_file *archivo, void*v){
    long total;
    long libre;
    long consumida;
    long compartida;
    long buffer;
    long swap;
    si_meminfo(&info);
    total=info.totalram*info.mem_unit;
    libre=info.freeram*info.mem_unit;
    compartida=info.sharedram*info.mem_unit;
    buffer=info.bufferram*info.mem_unit;
    swap=info.freeswap*info.mem_unit;
    consumida=total-libre-compartida-buffer-swap;
    
    seq_printf(archivo,"{\"Total\":\"%8li\",\n",total);
    seq_printf(archivo,"\"Consumida\":\"%8li\"}\n",consumida);
    
    return 0;
}

static int al_abrir(struct inode *inode,struct file *file){
    return single_open(file,escribir_archivo,NULL);
}

static struct proc_ops operaciones={
    .proc_open=al_abrir,
    .proc_read=seq_read
};

static int _insert(void)
{
    proc_create("Modulo de RAM",0,NULL,&operaciones);
    printk(KERN_INFO "Mensaje al insertar el modulo, Lab SO\n");
    
    return 0;
} 

static void _remove(void)
{
    remove_proc_entry("Modulo de RAM",NULL);
    printk(KERN_INFO "Mensaje al remover el modulo, Lab SO\n");
}

module_init(_insert);
module_exit(_remove);