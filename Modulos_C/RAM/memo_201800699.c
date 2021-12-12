//Header obligatorio de todos los modulos
#include <linux/module.h>
//Header para usar KERN_INFO
#include <linux/kernel.h>

//Header para los macros module_init y module_exit
#include <linux/init.h>
//Header necesario porque se usara proc_fs
#include <linux/proc_fs.h>
/* for copy_from_user */
#include <asm/uaccess.h>	
/* Header para usar la lib seq_file y manejar el archivo en /proc*/
#include <linux/seq_file.h>
#include <linux/hugetlb.h>

MODULE_LICENSE("GPL");
MODULE_DESCRIPTION("Modulo Memoria Ram");
MODULE_AUTHOR("Ronald Alexandro Castro Perez");

struct sysinfo inf;

//Funcion que se ejecutara cada vez que se lea el archivo con el comando CAT
static int escribir_archivo(struct seq_file *archivo, void *v)
{   
    long memoria_free;
    long memoria_total;
    long memoria_usada;
    long porcentaje;
    si_meminfo(&inf);

    memoria_total = inf.totalram * inf.mem_unit / 1048576;
    memoria_free = (inf.freeram + inf.bufferram) * inf.mem_unit/1048576;
    memoria_usada = memoria_total - memoria_free;
    porcentaje = 0;

    seq_printf(archivo, "{\"mem_total\":%ld,\n", memoria_total);
    seq_printf(archivo, "\"mem_free\":%ld,\n", memoria_free);
    seq_printf(archivo, "\"mem_used\":%ld,\n", memoria_usada);
    seq_printf(archivo, "\"percent\":%ld}", porcentaje);
    
    return 0;
}

//Funcion que se ejecutara cada vez que se lea el archivo con el comando CAT
static int al_abrir(struct inode *inode, struct file *file)
{
    return single_open(file, escribir_archivo, NULL);
}

//Si el kernel es 5.6 o mayor se usa la estructura proc_ops
static struct proc_ops operaciones =
{
    .proc_open = al_abrir,
    .proc_read = seq_read
};

/*Si el kernel es menor al 5.6 usan file_operations
static struct file_operations operaciones =
{
    .open = al_abrir,
    .read = seq_read
};
*/

//Funcion a ejecuta al insertar el modulo en el kernel con insmod
static int _insert(void)
{
    proc_create("memo_201800699", 0, NULL, &operaciones);
    printk(KERN_INFO "> Carnet: 20180699\n");
    return 0;
}

//Funcion a ejecuta al remover el modulo del kernel con rmmod
static void _remove(void)
{
    remove_proc_entry("memo_201800699", NULL);
    printk(KERN_INFO "Sistemas Operativos 1\n");
}

module_init(_insert);
module_exit(_remove);